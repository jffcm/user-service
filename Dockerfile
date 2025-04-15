FROM golang:1.24 as builder
WORKDIR /app
COPY ../go.mod ../go.sum ./
RUN go mod download
COPY ../ .
RUN go build -o user-service .

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/user-service .
EXPOSE 8080
CMD ["./user-service"]