FROM golang:1.23.4

WORKDIR /minesweeper

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -o minesweeper

CMD ["./minesweeper"]