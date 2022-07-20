package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	// "time"
	"html/template"
	"io/fs"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"

	"github.com/Pallinder/go-randomdata"

	"github.com/safanaj/svelte-go-only/pb"
	// "github.com/webview/webview"
)

var version, progname string

type CategoryServer struct {
	pb.UnimplementedCategoryServiceServer
}

func (_ *CategoryServer) Index(
	ctx context.Context, in *pb.IndexRequest) (*pb.Categories, error) {

	cat := pb.Categories{
		Categories: []*pb.Category{},
	}

	for i := 0; i < 1000; i++ {
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
}

func (h *rootHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	if h.grpcwebHandler.IsGrpcWebRequest(req) ||
		h.grpcwebHandler.IsAcceptableGrpcCorsRequest(req) ||
		contentType == "application/grpc-web+proto" {
		log.Printf("A content for GRPC-Web: %s %s %s", req.Proto, req.Method, req.URL.Path)
		h.grpcwebHandler.ServeHTTP(w, req)
		return
	}
	h.ginHandler.ServeHTTP(w, req)
}

//go:embed webui/dist
var distFS embed.FS
var assetsFS fs.FS
var templatesFS fs.FS

const distFSPrefix = "webui/dist"

var shutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

func main() {
	var err error

	assetsFS, err = fs.Sub(distFS, path.Join(distFSPrefix, "assets"))
	if err != nil {
		panic(err)
	}
	templatesFS, err = fs.Sub(distFS, path.Join(distFSPrefix, "templates"))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	wrappedServer := grpcweb.WrapServer(grpcServer)
	pb.RegisterCategoryServiceServer(grpcServer, &CategoryServer{})

	router := gin.Default()
	// if tmpl, err := template.ParseFS(templatesFS, "*.tmpl"); err != nil {
	// 	panic(err)
	// } else {
	// 	router.SetHTMLTemplate(tmpl)
	// }
	if html, err := template.ParseFS(distFS, path.Join(distFSPrefix, "*.html")); err != nil {
		panic(err)
	} else {
		router.SetHTMLTemplate(html)
	}

	// router.StaticFileFS("/icon", path.Join(distFSPrefix, "vite.svg"), http.FS(distFS))
	iconFileNames, err := assetsFS.(fs.GlobFS).Glob("vite.*.svg")
	if err != nil {
		panic(err)
	}
	if len(iconFileNames) < 1 {
		panic(fmt.Errorf("icon not found"))
	}

	router.StaticFS("/assets", http.FS(assetsFS))

	serveIndex := func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Svelte+ViteJS by Gin+GRPC-Web",
			"icon":  path.Join("assets", iconFileNames[0]),
		})
	}
	router.GET("/", serveIndex)
	router.HEAD("/", serveIndex)

	// Listen and serve on 0.0.0.0:8080
	// router.Run(":8080")
	mainCtx, mainCtxCancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithCancel(mainCtx)

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

	// test webview, avoid, it requires gtk3
	// {
	// 	wv := webview.New(true)
	// 	defer wv.Destroy()
	// 	wv.Navigate("http://localhost:8080/")
	// 	wv.Run()
	// }

	go func() {
		<-ctx.Done()
		srv.Shutdown(context.Background())
		log.Printf("Server shutdown done, going to close ...")
		mainCtxCancel()
	}()

	<-mainCtx.Done()
}
