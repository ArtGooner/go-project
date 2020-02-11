FROM golang:alpine
LABEL maintainer="Artsiom Ivaniutsenka <artstop92@gmail.com>"
RUN mkdir -p /app
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN chmod +x /app
RUN go build -o main ./
EXPOSE 50051/tcp
CMD ["./main"]