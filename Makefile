VERSION ?= 0.1.0
COMPONENT = svelte-go-only
FLAGS =
ENVVAR = CGO_ENABLED=1
GOOS ?= $(shell go env GOOS) #linux
GO ?= go
LDFLAGS ?= -s -w

SRCS = ./main.go
PROTO_SRCS = proto/*.proto

golang:
	@echo "--> Go Version"
	@$(GO) version

clean:
	rm -f main $(COMPONENT)
	rm -rf webui/dist webui/build pb webui/pb

clean-all: clean
	rm -rf webui/node_modules vendor

go-deps:
	mkdir -p $(CURDIR)/pb && echo "package pb" > $(CURDIR)/pb/doc.go
	$(GO) mod tidy -v && $(GO) mod vendor -v

go-deps-verify: go-deps
	$(GO) mod verify

proto/.built: $(PROTO_SRCS)
	$(MAKE) build-proto
	touch $@

build-proto:
	mkdir -p $(CURDIR)/pb $(CURDIR)/webui/pb
	cd $(CURDIR)/proto ; protoc *.proto \
		--go_out=paths=source_relative:$(CURDIR)/pb \
		--go-grpc_out=paths=source_relative:$(CURDIR)/pb \
		--js_out=import_style=commonjs,binary:$(CURDIR)/webui/pb \
		--ts_out=service=grpc-web,import_style=es6:$(CURDIR)/webui/pb

	for f in $(CURDIR)/webui/pb/*_pb.js; do cat $${f} | ./scripts/fix-js-for-es6.py | sponge $${f} ; done

webui/node_modules:
	cd webui ; yarn install

build-webui: webui/node_modules proto/.built
	cd webui ; yarn build

webui/build: proto/.built
	$(MAKE) build-webui
	touch $@

vendor:
	$(MAKE) go-deps
	touch $@

build-go: vendor webui/build proto/.built
	$(ENVVAR) GOOS=$(GOOS) $(GO) build -mod=vendor \
		-gcflags "-e" \
		-ldflags "$(LDFLAGS) -X main.version=${VERSION} -X main.progname=${COMPONENT}" \
		-v -o ${COMPONENT} $(SRCS)

build-static-go: vendor webui/build proto/.built
	@echo "--> Compiling the static binary"
	$(ENVVAR) GOARCH=amd64 GOOS=$(GOOS) $(GO) build -mod=vendor -a -tags netgo \
		-gcflags "-e" \
		-ldflags "$(LDFLAGS) -X main.version=${VERSION} -X main.progname=${COMPONENT}" \
		-v -o ${COMPONENT} $(SRCS)

build: build-go

start: build
	./${COMPONENT}

start-app: build
	./${COMPONENT} --as-app
