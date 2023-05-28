FROM golang:1.19-alpine

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .

ARG MAINDIR=rpc
RUN go build -ldflags="-s -w" -o /app ./$MAINDIR