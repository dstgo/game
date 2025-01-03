# basic info
app := wilson-service-game
module := github.com/dstgo/game/cmd/game
# meta info
author = dstgo
build_time := $(shell date +"%Y-%m-%dT%H:%M:%S")
git_version := $(shell git tag --sort=-version:refname | sed -n 1p)

# build info
mode := debug
output := $(shell pwd)/bin
host_os := $(shell go env GOHOSTOS)
host_arch := $(shell go env GOHOSTARCH)
os := $(host_os)
arch := $(host_arch)
ldflags := $(nullstring)

# reduce binary size at release mode
ifeq ($(mode), release)
	ldflags += -s -w
endif

# inject meta info
ifneq ($(app), $(nullstring))
	ldflags += -X main.AppName=$(app)
endif
ifneq ($(author), $(nullstring))
	ldflags += -X main.Author=$(author)
endif
ifneq ($(build_time), $(nullstring))
	ldflags += -X main.BuildTime=$(build_time)
endif
ifneq ($(git_version), $(nullstring))
	ldflags += -X main.Version=$(git_version)
endif

# binary extension
exe = $(nullstring)
ifeq ($(os), windows)
	exe = .exe
endif

.PHONY: build
build:
	# go lint
	go vet ./...

	# prepare target environment $(os)/$(arch)
	go env -w GOOS=$(os)
	go env -w GOARCH=$(arch)

	# build go module
	go build -trimpath \
		-ldflags="$(ldflags)" \
		-o $(output)/$(mode)/$(app)-$(os)-$(arch)/$(app)$(exe) \
		$(module)

	# resume host environment $(host_os)/$(host_arch)
	go env -w GOOS=$(host_os)
	go env -w GOARCH=$(host_arch)


api_dir := ./api
api_pb_files := $(shell find $(api_dir) -name "*.proto")
api_pb_gen_files := $(shell find $(api_dir) -type f \( -name "*.pb.go" -o -name "*.pb.*.go" \))

.PHONY: gen_pb
gen_pb:
	protoc --proto_path=./\
              --proto_path=./third_party \
              --go_out=paths=source_relative:.\
              --go-http_out=paths=source_relative:.\
              --go-grpc_out=paths=source_relative:.\
              --go-errors_out=paths=source_relative:.\
              --validate_out=lang=go,paths=source_relative:.\
              $(api_pb_files)

api_doc := $(api_dir)/v1

.PHONY: gen_doc
gen_doc:
	protoc --proto_path=./\
                  --proto_path=./third_party \
                  --openapi_out=fq_schema_naming=true,default_response=false:$(api_doc)\
                  $(api_pb_files)

.PHONY: rm_pb
rm_pb:
	rm -rf $(api_pb_gen_files)