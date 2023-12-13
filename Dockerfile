FROM golang:1.20

WORKDIR /go/src/go-fiber-api-docker

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]