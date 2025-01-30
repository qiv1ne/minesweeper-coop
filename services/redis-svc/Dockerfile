FROM golang

WORKDIR /usr/src/redis-svc

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/redis-svc ./main.go

CMD ["redis-svc"]
