FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    wget \
    curl \
    git \
    build-essential \
    ca-certificates \
    && wget https://golang.org/dl/go1.22.4.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.22.4.linux-amd64.tar.gz \
    && rm go1.22.4.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go

WORKDIR /app

COPY . .

RUN go build -o main ./cmd

EXPOSE 4000

CMD ["./main"]
