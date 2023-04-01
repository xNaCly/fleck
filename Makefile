FILES := $(shell find . -name "*.c")
dev build:
	./fleck

build:
	gcc $(FILES)\
		-O3 \
		-fdiagnostics-color=always  \
		-Wall \
		-Wpedantic \
		-std=c11 \
		-Wextra \
		-Werror \
		-Wshadow \
		-Wundef \
		-fno-common \
		-o fleck

