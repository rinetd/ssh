COMPILE_TIME = $(shell date +"%Y-%M-%d %H:%M:%S")
BUILD=`date +%FT%T%z`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS_f1=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD} -X main.Entry=f1"
LDFLAGS_f2=-ldflags "-w -s -X main.Version=${VERSION} -X main.Build=${BUILD} -X main.Entry=f2"


name:=getdata
# make ver=release
ifeq ($(ver), debug)
	CXXFLAGS = -c -g -Ddebug
else
	CXXFLAGS = -c -O3 
endif

.PHONY: all
all: windows

.PHONY: windows-dependencies
windows-dependencies:
	go get github.com/josephspurrier/goversioninfo/cmd/goversioninfo

.PHONY: embed-assets
embed-assets:
	@# go get github.com/jteeuwen/go-bindata/...
	@# go-bindata ./logos/$(name)_headert.png ./webpage.html

.PHONY: linux
linux: *.go embed-assets
	GOOS=linux GOARCH=amd64 go build -o output/$(name)_linux_64bit
	GOOS=linux GOARCH=386 go build -o output/$(name)_linux_32bit
	strip output/$(name)_linux_*

.PHONY: windows 
windows: *.go windows-dependencies embed-assets
	goversioninfo -icon=rc/icon.ico -manifest=rc/manifest.exe.manifest rc/versioninfo.json
	@- rm output/*.exe 
	CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags="-s -w " -o output/$(name)_`date +%m-%d`.exe
	@# - rm output/$(name)_windows_64bit.exe
	@# - CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui -linkmode internal" -o output/deploy_windows_64bit.exe
	rm resource.syso

.PHONY: osx
osx: *.go embed-assets
	GOOS=darwin GOARCH=amd64 go build -o output/$(name)_osx_64bit
	GOOS=darwin GOARCH=386 go build -o output/$(name)_osx_32bit


.PHONY: clean
clean:
	rm -rf output/*