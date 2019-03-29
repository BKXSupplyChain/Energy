FROM golang:1.7.3 AS builder

RUN go get -d -v github.com/globalsign/mgo
RUN go get -d -v github.com/micro/go-config
RUN go get -d -v github.com/ethereum/go-ethereum/crypto

ADD . /go/src/github.com/BKXSupplyChain/Energy
WORKDIR /go/src/github.com/BKXSupplyChain/Energy/backend
RUN CGO_ENABLED=0 GOOS=linux go build -a -o backend .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/BKXSupplyChain/Energy/backend/backend .
COPY --from=builder /go/src/github.com/BKXSupplyChain/Energy/backend/conf.json .
COPY --from=builder /go/src/github.com/BKXSupplyChain/Energy/backend/web/static ./web/static/
CMD ["./backend"]