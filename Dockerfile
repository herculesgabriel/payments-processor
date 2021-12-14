FROM golang:1.17

WORKDIR /go/src

RUN apt-get update && apt-get install build-essential librdkafka-dev -y

RUN apt-get install curl && \
  curl https://raw.githubusercontent.com/eficode/wait-for/v2.1.3/wait-for --output /usr/bin/wait-for && \
  chmod +x /usr/bin/wait-for

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
