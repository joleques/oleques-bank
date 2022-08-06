FROM golang:1.17.0-alpine3.14 as builder

WORKDIR /go/src/github.com/joleques/oleques-bank/transfer

RUN apk update && apk add build-base

COPY . .

RUN go mod download

RUN go get github.com/golang/mock/gomock@v1.6.0
RUN go install github.com/golang/mock/mockgen@v1.6.0
RUN ./../../../../../go/bin/mockgen -source=src/domain/ports.go -destination=pkg/domain/mock/trasnserPorts.go -package=mock

RUN go test ./...

RUN go build -ldflags "-s -w" src/main.go

FROM alpine:3.14

WORKDIR /app

COPY --from=builder /go/src/github.com/joleques/oleques-bank/transfer/main .

ENV COMMIT=@@commit
ADD k8s k8s

CMD [ "./main" ]
