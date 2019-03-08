FROM golang:1.7.3 AS builder

RUN go get -d -v github.com/globalsign/mgo
RUN go get -d -v github.com/micro/go-config

ADD . /go/src/github.com/BKXSupplyChain/Energy
WORKDIR /go/src/github.com/BKXSupplyChain/Energy/sensors
RUN CGO_ENABLED=0 GOOS=linux go build -a -o sensors .

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/BKXSupplyChain/Energy/sensors/sensors .
COPY --from=builder /go/src/github.com/BKXSupplyChain/Energy/sensors/config.json .
CMD ["./sensors"]