FROM golang:1.18
WORKDIR /go/src/go-fiber-api-docker
COPY . .

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]