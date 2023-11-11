export GO111MODULE=on
.PHONY: clean fmt generate lint test unit vendor

ARCH ?= amd64
OS ?= linux
ALL_ARCH := amd64 arm arm64
DOCKER_ARCH := "amd64" "arm v7" "arm64 v8"
BINS := $(addprefix bin/$(ARCH)/,terraform-provider-saladcloud)
PROJECT := terraform-provider-saladcloud
PKG := github.com/squat/$(PROJECT)

TAG := $(shell git describe --abbrev=0 --tags HEAD 2>/dev/null)
COMMIT := $(shell git rev-parse HEAD)
VERSION := $(COMMIT)
ifneq ($(TAG),)
    ifeq ($(COMMIT), $(shell git rev-list -n1 $(TAG)))
        VERSION := $(TAG)
    endif
endif
DIRTY := $(shell test -z "$$(git diff --shortstat 2>/dev/null)" || echo -dirty)
VERSION := $(VERSION)$(DIRTY)
LD_FLAGS := -ldflags "-X main.Version=$(VERSION) -extldflags -static"
GO_FILES := $(shell find . -name '*.go' -not -path './vendor/*')
GO_PKGS := $(shell go list ./... | grep -v "$(PKG)/vendor")
EXAMPLES := $(shell find examples -type f -name '*.tf')
SPEAKEASY_FILES := $(shell cat files.gen)
DOCS := $(shell find docs -type f -name '*.md')
ifeq ($(DOCS),)
DOCS := docs/index.md
endif
GENERATED := $(SPEAKEASY_FILES) $(DOCS) files.gen

STATICCHECK_BINARY := bin/staticcheck
YQ_BINARY := bin/yq
SPEAKEASY_BINARY := bin/speakeasy
GOJSONTOYAML_BINARY := bin/gojsontoyaml

GO_VERSION ?= 1.21.3
BUILD_IMAGE ?= golang:$(GO_VERSION)-alpine

build: $(BINS)

build-%:
	@$(MAKE) --no-print-directory ARCH=$* build

all-build: $(addprefix build-, $(ALL_ARCH))

CONTAINERIZE_BUILD ?= true
BUILD_PREFIX :=
BUILD_SUFIX :=
ifeq ($(CONTAINERIZE_BUILD), true)
	BUILD_PREFIX := docker run --rm \
	    -u $$(id -u):$$(id -g) \
	    -v $$(pwd):/src \
	    -w /src \
	    --entrypoint '' \
	    $(BUILD_IMAGE) \
	    /bin/sh -c ' \
	        GOCACHE=$$(pwd)/.cache
	BUILD_SUFIX := '
endif

$(BINS): $(GO_FILES) go.mod | files.gen
	@mkdir -p bin/$(ARCH)
	@echo "building: $@"
	@$(BUILD_PREFIX) \
	        GOARCH=$(ARCH) \
	        GOOS=linux \
		CGO_ENABLED=0 \
		go build -o $@ \
		    $(LD_FLAGS) \
		    . \
	$(BUILD_SUFIX)

fmt:
	@echo $(GO_PKGS)
	gofmt -w -s $(GO_FILES)

lint: $(STATICCHECK_BINARY)
	@echo 'go vet $(GO_PKGS)'
	@vet_res=$$(GO111MODULE=on go vet $(GO_PKGS) 2>&1); if [ -n "$$vet_res" ]; then \
		echo ""; \
		echo "Go vet found issues. Please check the reported issues"; \
		echo "and fix them if necessary before submitting the code for review:"; \
		echo "$$vet_res"; \
		exit 1; \
	fi
	@echo '$(STATICCHECK_BINARY) $(GO_PKGS)'
	@lint_res=$$($(STATICCHECK_BINARY) $(GO_PKGS)); if [ -n "$$lint_res" ]; then \
		echo ""; \
		echo "Staticcheck found style issues. Please check the reported issues"; \
		echo "and fix them if necessary before submitting the code for review:"; \
		echo "$$lint_res"; \
		exit 1; \
	fi
	@echo 'gofmt -d -s $(GO_FILES)'
	@fmt_res=$$(gofmt -d -s $(GO_FILES)); if [ -n "$$fmt_res" ]; then \
		echo ""; \
		echo "Gofmt found style issues. Please check the reported issues"; \
		echo "and fix them if necessary before submitting the code for review:"; \
		echo "$$fmt_res"; \
		exit 1; \
	fi

unit:
	go test --race ./...

test: lint unit

bin-clean:
	rm -rf bin

vendor:
	go mod tidy
	go get toolchain@go$(GO_VERSION)

$(STATICCHECK_BINARY):
	go build -o $@ honnef.co/go/tools/cmd/staticcheck

$(YQ_BINARY):
	go build -o $@ github.com/mikefarah/yq/v4

$(SPEAKEASY_BINARY):
	cd $(@D) && curl https://github.com/speakeasy-api/speakeasy/releases/download/v1.119.0/speakeasy_$(OS)_$(ARCH).zip -L -o speakeasy.zip && unzip -o speakeasy.zip $(@F) && rm speakeasy.zip ; chmod +x $(@F)

$(GOJSONTOYAML_BINARY):
	go build -o $@ github.com/brancz/gojsontoyaml

saladcloud.json:
	curl https://docs.salad.com/reference/saladcloud-openapi-spec?json=on | jq .doc.body -r | head --lines=-1 | tail --lines=+2

saladcloud.yaml: saladcloud.json $(GOJSONTOYAML_BINARY) $(YQ_BINARY) patch.sh
	cat $< | $(GOJSONTOYAML_BINARY) | sh patch.sh > $@

$(SPEAKEASY_FILES) files.gen &: saladcloud.yaml $(SPEAKEASY_BINARY)
	$(SPEAKEASY_BINARY) generate sdk --lang terraform --schema $< --out .
	# Add tools.
	sed -i '\|// Documentation generation|i _ "honnef.co/go/tools/cmd/staticcheck"' tools/tools.go
	sed -i '\|// Documentation generation|i _ "github.com/mikefarah/yq/v4"' tools/tools.go
	sed -i '\|// Documentation generation|i _ "github.com/brancz/gojsontoyaml"' tools/tools.go
	grep "some_saladcloud_api_key" --files-without-match $$(grep "# Configuration options" --recursive --exclude-dir bin --files-with-matches .) | xargs sed -i '/^\s*# Configuration options/a \ \ api_key_auth = "some_saladcloud_api_key"'
	$(MAKE) fmt
	$(MAKE) vendor

$(DOCS) &: $(EXAMPLES) $(SPEAKEASY_FILES) | files.gen
	go generate
	sed -i 's/saladcloud Provider/SaladCloud Provider/' docs/index.md

generate: $(GENERATED)
