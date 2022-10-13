FROM --platform=linux/arm64 golang:1.17 AS build

ENV GO111MODULE=on

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y git

WORKDIR /go/src/github.com/steve-shop/

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o main .

FROM alpine
COPY --from=build /go/src/github.com/steve-shop/main go/bin/steve-shop/

EXPOSE 9000
WORKDIR go/bin/steve-shop/

CMD ["./main"]