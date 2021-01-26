.PHONY: clean

zigandgotest: main.go libadd.so
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="$(PWD)/bin/zcc" CXX="$(PWD)/bin/zxx" go build --tags extended

libadd.so: add.zig
	zig build-lib -dynamic add.zig

clean:
	rm -rf hello *.a *.o *.so zig-cache
