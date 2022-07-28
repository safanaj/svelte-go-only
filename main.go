package main

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	"github.com/Pallinder/go-randomdata"
	"github.com/google/uuid"
	"github.com/prometheus/procfs"

	"github.com/safanaj/svelte-go-only/pb"
	// "github.com/webview/webview"
)

var version, progname string

var procFS procfs.FS

type ControlMsgServer struct {
	pb.UnimplementedControlMsgServiceServer

	refreshCh   chan string
	uuid2stream map[string]pb.ControlMsgService_ControlServer
}

func (s *ControlMsgServer) startRefresher(ctx context.Context) {
	s.refreshCh = make(chan string)
	s.uuid2stream = make(map[string]pb.ControlMsgService_ControlServer)
	sendControlMsg := func(stream pb.ControlMsgService_ControlServer, t time.Time) error {
		la, _ := procFS.LoadAvg()
		mi, _ := procFS.Meminfo()
		mu := "unknown"
		if mi.MemTotal != nil && mi.MemAvailable != nil {
			mu = fmt.Sprintf("%d", (*mi.MemTotal)-(*mi.MemAvailable))
		}
		cmsg := &pb.ControlMsg{
			Date:     t.Format(time.RFC850),
			CpuUsage: fmt.Sprintf("%g", la.Load1),
			MemUsage: mu,
		}
		return stream.Send(cmsg)
	}

	ticker := time.NewTicker(10 * time.Second)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case t := <-ticker.C:
				toDel := []string{}
				for ruuid, stream := range s.uuid2stream {
					if err := sendControlMsg(stream, t); err != nil {
						toDel = append(toDel, ruuid)
					}
				}
				for _, x := range toDel {
					delete(s.uuid2stream, x)
				}
			case ruuid := <-s.refreshCh:
				if stream, ok := s.uuid2stream[ruuid]; ok {
					if err := sendControlMsg(stream, time.Now()); err != nil {
						delete(s.uuid2stream, ruuid)
					}
				} else {
					fmt.Printf("Wrong uuid: %s\n", ruuid)
				}
			}
		}
	}()
}

func (s *ControlMsgServer) Refresh(
	ctx context.Context, in *pb.ControlMsgEmpty) (*pb.ControlMsgEmpty, error) {
	ruuid, isOk := ctx.Value(idCtxKey).(string)
	if isOk {
		s.refreshCh <- ruuid
	}
	return in, nil
}

func (s *ControlMsgServer) Control(stream pb.ControlMsgService_ControlServer) error {
	ruuid, isOk := stream.Context().Value(idCtxKey).(string)
	if isOk {
		// fmt.Printf("UUID: %s\n", ruuid)
		s.uuid2stream[ruuid] = stream
	} else {
		fmt.Printf("Not UUID in stream Context: %v\n", stream.Context().Value(idCtxKey))
	}
	select {
	case <-stream.Context().Done():
		return nil
	}
}

type CategoryServer struct {
	pb.UnimplementedCategoryServiceServer
}

func (_ *CategoryServer) Index(
	ctx context.Context, in *pb.IndexRequest) (*pb.Categories, error) {

	cat := pb.Categories{
		Categories: []*pb.Category{},
	}

	for i := 0; i < int(in.Number); i++ {
		category := &pb.Category{
			Id: fmt.Sprintf("%v", i),
		}
		if in.Kind == pb.IndexRequest_COUNTRY {
			category.Name = randomdata.Country(randomdata.ThreeCharCountry)
		} else if in.Kind == pb.IndexRequest_CITY {
			category.Name = randomdata.City()
		} else if in.Kind == pb.IndexRequest_ADJECTIVE {
			category.Name = randomdata.Adjective()
		} else if in.Kind == pb.IndexRequest_EMAIL {
			category.Name = randomdata.Email()
		} else if in.Kind == pb.IndexRequest_CURRENCY {
			category.Name = randomdata.Currency()
		} else {
			category.Name = randomdata.City()
		}
		cat.Categories = append(cat.Categories, category)
	}

	return &cat, nil
}

type rootHandler struct {
	ginHandler     *gin.Engine
	grpcwebHandler *grpcweb.WrappedGrpcServer

	counter int
}

type ctxKey struct{ name string }

const idKey = "ruuid"

var idCtxKey = &ctxKey{name: "id"}

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	connection := req.Header.Get("Connection")
	upgrade := req.Header.Get("Upgrade")
	wsProtocol := req.Header.Get("Sec-Websocket-Protocol")

	var ruuid string
	cookie, err := req.Cookie(idKey)
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{Name: idKey, Value: fmt.Sprintf("%s", uuid.New())}
	}
	http.SetCookie(w, cookie)
	rctx := context.WithValue(req.Context(), idCtxKey, cookie.Value)
	if h.grpcwebHandler.IsGrpcWebRequest(req) ||
		h.grpcwebHandler.IsAcceptableGrpcCorsRequest(req) ||
		contentType == "application/grpc-web+proto" ||
		(connection == "Upgrade" && upgrade == "websocket" && wsProtocol == "grpc-websockets") {
		log.Printf("A content for GRPC-Web: %s %s %s", req.Proto, req.Method, req.URL.Path)
		h.grpcwebHandler.ServeHTTP(w, req.WithContext(rctx))
		return
	}
	h.ginHandler.ServeHTTP(w, req.WithContext(rctx))
}

//go:embed all:webui/build
var distFS embed.FS

// //go:embed webui/build/_app
// var appFS embed.FS
var appFS fs.FS
var assetsFS fs.FS

// var templatesFS fs.FS

const distFSPrefix = "webui/build"
const appFSPrefix = "webui/build/_app"
const assetsFSPrefix = "immutable/assets"

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

func main() {
	var err error

	procFS, err = procfs.NewDefaultFS()
	if err != nil {
		panic(err)
	}

	appFS, err = fs.Sub(distFS, path.Join(distFSPrefix, "_app"))
	if err != nil {
		panic(err)
	}

	assetsFS, err = fs.Sub(appFS, assetsFSPrefix)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	wrappedServer := grpcweb.WrapServer(grpcServer, grpcweb.WithWebsockets(true))

	mainCtx, mainCtxCancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithCancel(mainCtx)

	cMsgSrv := &ControlMsgServer{}
	pb.RegisterControlMsgServiceServer(grpcServer, cMsgSrv)
	pb.RegisterCategoryServiceServer(grpcServer, &CategoryServer{})

	cMsgSrv.startRefresher(ctx)

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	if html, err := template.ParseFS(distFS, path.Join(distFSPrefix, "*.html")); err != nil {
		panic(err)
	} else {
		router.SetHTMLTemplate(html)
	}

	iconFileNames, err := assetsFS.(fs.GlobFS).Glob("vite-*.svg")
	if err != nil {
		panic(err)
	}
	if len(iconFileNames) < 1 {
		panic(fmt.Errorf("icon not found"))
	}

	router.StaticFS("/_app", http.FS(appFS))

	serveHTML := func(page string) func(*gin.Context) {
		return func(c *gin.Context) {
			c.HTML(http.StatusOK, page, gin.H{
				"title": "Svelte+ViteJS by Gin+GRPC-Web",
				"icon":  path.Join("/_app/immutable/assets", iconFileNames[0]),
			})
		}
	}

	router.GET("/", serveHTML("index.html"))
	router.HEAD("/", serveHTML("index.html"))
	router.GET("/about", serveHTML("about.html"))
	router.HEAD("/about", serveHTML("about.html"))

	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		cancel()
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	srv := &http.Server{Addr: ":8080", Handler: &rootHandler{ginHandler: router, grpcwebHandler: wrappedServer}}
	go srv.ListenAndServe()

	go func() {
		<-ctx.Done()
		srv.Shutdown(context.Background())
		log.Printf("Server shutdown done, going to close ...")
		mainCtxCancel()
	}()

	<-mainCtx.Done()
}
