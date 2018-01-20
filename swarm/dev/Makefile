.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/one-currency/one-coin/cmd/swarm
	go build -o bin/geth     github.com/one-currency/one-coin/cmd/geth
	go build -o bin/bootnode github.com/one-currency/one-coin/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/one-currency/one-coin/swarm/...
