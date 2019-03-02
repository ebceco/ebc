.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/ebceco/ebc/cmd/swarm
	go build -o bin/geth     github.com/ebceco/ebc/cmd/geth
	go build -o bin/bootnode github.com/ebceco/ebc/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/ebceco/ebc/swarm/...
