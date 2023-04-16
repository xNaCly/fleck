ver := "0.1.0"
opt := ""

dev: pre
	mkdir -p ./out/dev
	go build -o ./out/dev/

release: pre
	mkdir -p ./out/release
	go build -o ./out/release/

pre:
	mkdir -p ./out
