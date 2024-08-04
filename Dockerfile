ARG GO_VERSION=1.22.4 \
    ALPINE_VERSION=3.20

FROM golang:${GO_VERSION}-alpine${ALPINE_VERSION} AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

# Copy the source code
COPY main.go .
COPY cmd/* cmd/

RUN CGO_ENABLED=0 go build -o kafka-client .

FROM alpine:${ALPINE_VERSION}

WORKDIR /app

COPY --from=builder /app/kafka-client .

CMD ["./kafka-client"]
