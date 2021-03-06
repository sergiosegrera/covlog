FROM golang:latest AS builder
WORKDIR /service
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o service.o ./cmd/main.go

FROM alpine:latest
WORKDIR /service
COPY --from=builder /service/service.o /service/

ENTRYPOINT ./service.o