FROM golang:1.22.4 AS builder

WORKDIR /app

COPY . .

RUN go build -o main ./cmd

FROM gcr.io/distroless/base-debian11

WORKDIR /app

COPY --from=builder /app/main /app/

EXPOSE 4000

CMD ["./main"]
