.PHONY: clean run

zigandgotest: main.go lib/libadd.so
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="$(PWD)/bin/zcc" CXX="$(PWD)/bin/zxx" go build --tags extended

lib/libadd.so: add.zig
	zig build-lib -dynamic -rdynamic add.zig
	mkdir -p lib
	mv libadd.so lib

clean:
	rm -rf zigandgotest zig-cache lib/*.so

run: zigandgotest
	@ldd zigandgotest
	@./run.sh # this works after running "go build"
