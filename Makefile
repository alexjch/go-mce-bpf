libbpf:
	scripts/build_libbpf.sh

libbpf-clean:
	rm -rf ./root ./build

vmlinux.h:
	bpftool btf dump file /sys/kernel/btf/vmlinux format c > bpf/vmlinux.h

build: clean
	clang -g -O2 -target bpf -c bpf/mce_log.c -o mce_log.o

build-go:
	CC=gcc CGO_CFLAGS="-I ./root/usr/include" CGO_LDFLAGS="./root/usr/lib64/libbpf.a" go build -mod vendor -o mce_log cmd/mce_log/main.go

clean:
	rm mce_log*