To repro:

```
On a 24.04 ubuntu (GCE: ubuntu-os-cloud image, n2-custom-24-32768)

docker run --security-opt seccomp:unconfined -it

export PATH=/usr/local/go/bin:~/go/bin:$PATH
go install github.com/cockroachdb/stress@latest

git clone https://github.com/RaduBerinde/msan-repro.git
cd msan-repro

CC=clang go test -msan . --exec 'stress -p 40' -v
```

If it doesn't repro, interrupt and run again.

