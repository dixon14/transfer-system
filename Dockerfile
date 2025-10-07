# Build stage
FROM golang:1.24-alpine AS builder

WORKDIR /
COPY . .
RUN apk add g++ make
RUN make build


FROM alpine:latest
WORKDIR /
COPY --from=builder bin/transfer-system /app/transfer-system
RUN chmod +x /app/transfer-system

ENTRYPOINT [ "/app/transfer-system" ]