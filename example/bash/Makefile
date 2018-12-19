PROJECT = transfer
VERSION = $(shell git describe)
BRANCH = $(shell git name-rev --name-only HEAD)
VERSION ="$(echo $TRAVIS_TAG)"
ARCNAME = $(PROJECT)-$(VERSION)-$(GOOS)-$(GOARCH)
ARTIFACTS_DIR=artifact


# ZIP="zip -9 -r"
# TAR="tar cvf"
# GZIP="gzip -9"
# GZIP_7Z="7za a -tgzip -mx9"
# ZIP_7Z="7za a -tzip -mx9"
ifndef ZIP
	ifneq ($(shell which zip 2>/dev/null),)
		ZIP := zip -9
	endif
	ifneq ($(shell which 7z 2>/dev/null),)
		ZIP := 7z a -tzip -mx=9
	endif
	ifneq ($(shell which 7za 2>/dev/null),)
		ZIP := 7za a -tzip -mx=9
	endif
    ifndef ZIP
        $(warning "No zip / 7z / 7za in ($(PATH))")
        ZIP := : zip_not_found
    endif
endif

.PHONY: all release release-all
all: release-all


.PHONY: windows-dependencies
windows-dependencies:
	@ go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo

.PHONY: embed-assets
embed-assets:
	@ GO111MODULE=off go get -u github.com/jteeuwen/go-bindata/...
	@# go-bindata ./logos/microBadger_headert.png ./webpage.html



release:
	env CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build $(FLAGS) -ldflags='-s -w -X github.com/rinetd/transfer/version.Version=$(VERSION)' -o build/$(GOOS)-$(GOARCH)/$(PROJECT)$(EXT)
	- tar czf build/$(ARCNAME).tar.gz -C build/$(GOOS)-$(GOARCH)/ $(PROJECT)$(EXT)
	- tar czf build/$(ARCNAME).tar.gz -C build/$(GOOS)-$(GOARCH)/ $(PROJECT)$(EXT)
ifeq ($(OS),Windows_NT)
	cd dist; \
	rm -f $(ARCNAME)-*-Win.zip; \
	$(ZIP) $(ARCNAME)-$(VERSTR)-Win.zip build/$(GOOS)-$(GOARCH)/ $(PROJECT)$(EXT);
	$(ZIP) dist/$(ARCNAME)-$(VERSTR)-Win.zip  -r0 images/*.png; 
endif

release-all:
	# -@$(MAKE) release GOOS=darwin GOARCH=amd64
	# -@$(MAKE) release GOOS=linux GOARCH=386
	# -@$(MAKE) release GOOS=linux GOARCH=amd64
	-@$(MAKE) release GOOS=windows GOARCH=386 EXT=.exe
	# -@$(MAKE) release GOOS=windows GOARCH=amd64 EXT=.exe

.PHONY: windows 
windows: windows-dependencies embed-assets
	goversioninfo -icon=rc/icon.ico -manifest=rc/manifest.exe.manifest rc/versioninfo.json
	@- rm binaries/*.exe 
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -x -ldflags="-s -w " -o binaries/deploy`date +%m%d`.exe
	@# - rm binaries/deploy_windows_64bit.exe
	@# - CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui -linkmode internal" -o binaries/deploy_windows_64bit.exe
	rm resource.syso

platform:
	# @$(MAKE) releaseGOOS=js GOARCH=wasm
	# @$(MAKE) release GOOS=windows GOARCH=386 FLAGS='-ldflags="-H=windowsgui"' EXE=.exe
	# @$(MAKE) release GOOS=windows GOARCH=amd64 FLAGS='-ldflags="-H=windowsgui"' EXE=.exe
	# @$(MAKE) release GOOS=linux GOARCH=arm
	# @$(MAKE) release GOOS=linux GOARCH=arm64
	# @$(MAKE) release GOOS=linux GOARCH=mips
	# @$(MAKE) release GOOS=linux GOARCH=mips64
	# @$(MAKE) release GOOS=linux GOARCH=mips64le
	# @$(MAKE) release GOOS=linux GOARCH=mipsle
	# @$(MAKE) release GOOS=linux GOARCH=ppc64
	# @$(MAKE) release GOOS=linux GOARCH=ppc64le
	# @$(MAKE) release GOOS=linux GOARCH=s390x
	# @$(MAKE) release GOOS=android GOARCH=386
	# @$(MAKE) release GOOS=android GOARCH=amd64
	# @$(MAKE) release GOOS=android GOARCH=arm
	# @$(MAKE) release GOOS=android GOARCH=arm64
	# @$(MAKE) release GOOS=darwin GOARCH=386
	# @$(MAKE) release GOOS=darwin GOARCH=arm
	# @$(MAKE) release GOOS=darwin GOARCH=arm64
	# @$(MAKE) release GOOS=dragonfly GOARCH=amd64
	# @$(MAKE) release GOOS=freebsd GOARCH=386
	# @$(MAKE) release GOOS=freebsd GOARCH=amd64
	# @$(MAKE) release GOOS=freebsd GOARCH=arm
	# @$(MAKE) release GOOS=nacl GOARCH=386
	# @$(MAKE) release GOOS=nacl GOARCH=amd64p32
	# @$(MAKE) release GOOS=nacl GOARCH=arm
	# @$(MAKE) release GOOS=netbsd GOARCH=386
	# @$(MAKE) release GOOS=netbsd GOARCH=amd64
	# @$(MAKE) release GOOS=netbsd GOARCH=arm
	# @$(MAKE) release GOOS=openbsd GOARCH=386
	# @$(MAKE) release GOOS=openbsd GOARCH=amd64
	# @$(MAKE) release GOOS=openbsd GOARCH=arm
	# @$(MAKE) release GOOS=plan9 GOARCH=386
	# @$(MAKE) release GOOS=plan9 GOARCH=amd64
	# @$(MAKE) release GOOS=plan9 GOARCH=arm
	# @$(MAKE) release GOOS=solaris GOARCH=amd64


.PHONY: build
build:
	go get ./...

.PHONY: test
test:
	go get -t ./...
	go test -v ...

build-image:
	docker build -t rientd/transfer .

clean:
	rm -rf build dist
