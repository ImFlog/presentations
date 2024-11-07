FROM golang:1.23 AS builder

ENV CGO_ENABLED=0

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

COPY . .
RUN go build -o /app/bin .

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/bin /app/bin/app

EXPOSE 8080

ENTRYPOINT ["/app/bin/app"]
