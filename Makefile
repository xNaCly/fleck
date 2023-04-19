ver := 0.0.1-alpha
build_at := $(shell date +"%Y-%m-%dT%H:%M:%S%z")
build_by := $(shell git config --global user.name)-$(shell git config --global user.email)
feat := watch.1
opt := -ldflags="-w -s -X main.VERSION=$(ver)+$(feat) -X main.BUILD_AT=$(build_at) -X main.BUILD_BY=$(build_by)"

release: build_linux build_darwin build_windows

build_linux: pre
	env = CGO_ENABLED=0 GOOS=linux GOARCH=386
	$(env) go build -o ./out/linux/fleck_$(ver)+$(feat)_linux-i386 $(opt) 

	env = CGO_ENABLED=0 GOOS=linux GOARCH=386
	$(env) go build -tags=bare -o ./out/linux/fleck-bare_$(ver)+$(feat)_linux-i386 $(opt) 

	env = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	$(env) go build  -o ./out/linux/fleck_$(ver)+$(feat)_linux-x86_64 $(opt)

	env = CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	$(env) go build -tags=bare -o ./out/linux/fleck-bare_$(ver)+$(feat)_linux-x86_64 $(opt) 

	env = CGO_ENABLED=0 GOOS=linux GOARCH=arm
	$(env) go build  -o ./out/linux/fleck_$(ver)+$(feat)_linux-arm $(opt)

	env = CGO_ENABLED=0 GOOS=linux GOARCH=arm
	$(env) go build -tags=bare -o ./out/linux/fleck-bare_$(ver)+$(feat)_linux-arm $(opt) 

	env = CGO_ENABLED=0 GOOS=linux GOARCH=arm64
	$(env) go build  -o ./out/linux/fleck_$(ver)+$(feat)_linux-arm64 $(opt)

	env = CGO_ENABLED=0 GOOS=linux GOARCH=arm64
	$(env) go build -tags=bare -o ./out/linux/fleck-bare_$(ver)+$(feat)_linux-arm64 $(opt) 

build_windows: pre
	env = CGO_ENABLED=0 GOOS=windows GOARCH=386
	$(env) go build  -o ./out/windows/fleck_$(ver)+$(feat)_windows-i386.exe $(opt)

	env = CGO_ENABLED=0 GOOS=windows GOARCH=386
	$(env) go build -tags=bare -o ./out/windows/fleck-bare_$(ver)+$(feat)_windows-i386.exe $(opt) 

	env = CGO_ENABLED=0 GOOS=windows GOARCH=amd64
	$(env) go build  -o ./out/windows/fleck_$(ver)+$(feat)_windows-amd64.exe $(opt)

	env = CGO_ENABLED=0 GOOS=windows GOARCH=amd64
	$(env) go build -tags=bare -o ./out/windows/fleck-bare_$(ver)+$(feat)_windows-amd64.exe $(opt) 

	env = CGO_ENABLED=0 GOOS=windows GOARCH=arm
	$(env) go build  -o ./out/windows/fleck_$(ver)+$(feat)_windows-arm.exe $(opt)

	env = CGO_ENABLED=0 GOOS=windows GOARCH=arm
	$(env) go build -tags=bare -o ./out/windows/fleck-bare_$(ver)+$(feat)_windows-arm.exe $(opt) 

	env = CGO_ENABLED=0 GOOS=windows GOARCH=arm64
	$(env) go build  -o ./out/windows/fleck_$(ver)+$(feat)_windows-arm64.exe $(opt)

	env = CGO_ENABLED=0 GOOS=windows GOARCH=arm64
	$(env) go build -tags=bare -o ./out/windows/fleck-bare_$(ver)+$(feat)_windows-arm64.exe $(opt) 

build_darwin: pre
	env = CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
	$(env) go build  -o ./out/darwin/fleck_$(ver)+$(feat)_darwin-amd64 $(opt)

	env = CGO_ENABLED=0 GOOS=darwin GOARCH=amd64
	$(env) go build -tags=bare -o ./out/darwin/fleck-bare_$(ver)+$(feat)_darwin-amd64 $(opt) 

	env = CGO_ENABLED=0 GOOS=darwin GOARCH=arm64
	$(env) go build  -o ./out/darwin/fleck_$(ver)+$(feat)_darwin-arm64 $(opt)

	env = CGO_ENABLED=0 GOOS=darwin GOARCH=arm64
	$(env) go build -tags=bare -o ./out/darwin/fleck-bare_$(ver)+$(feat)_darwin-arm64 $(opt) 

pre: clean
	mkdir -p ./out
	mkdir -p ./out/darwin
	mkdir -p ./out/linux
	mkdir -p ./out/windows

clean:
	rm -fr ./out
