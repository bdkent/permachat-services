# bkent/permachat-identity-service

FROM golang:1.11.4

RUN apt-get update
RUN apt-get install sudo

RUN mkdir -p /go/src/github.com/bdkent
RUN git clone https://github.com/bdkent/permachat-services.git /go/src/github.com/bdkent/permachat-services
WORKDIR /go/src/github.com/bdkent/permachat-services/internal/app/identity
RUN go get -d -v ./...

WORKDIR /go/src/github.com/ethereum/go-ethereum
RUN go build -o /go/bin/abigen ./cmd/abigen

RUN git clone --recursive https://github.com/ethereum/solidity.git /tmp/solidity
WORKDIR /tmp/solidity
RUN git checkout v0.5.2
RUN ./scripts/install_deps.sh
RUN ./scripts/build.sh

RUN git clone https://github.com/bdkent/permachat.git /tmp/permachat
RUN  /go/bin/abigen --sol=/tmp/permachat/contracts/PermaChat.sol --pkg=main --out=/go/src/github.com/bdkent/permachat-services/internal/app/identity/PermaChatContract.go

WORKDIR /go/src/github.com/bdkent/permachat-services/internal/app/identity
RUN go build -o /go/bin/identity

RUN mkdir -p /etc/permachat
RUN touch /etc/permachat/config.toml

CMD /go/bin/identity
