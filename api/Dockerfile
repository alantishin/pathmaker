FROM golang:alpine3.16 as dev

WORKDIR /home/app

COPY . .

ENV HTTP_PORT=8080

RUN go mod download
RUN go install github.com/cosmtrek/air@v1.40.4

CMD air
