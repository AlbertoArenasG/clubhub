FROM golang:latest
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app ./cmd/server/main.go
EXPOSE 3000

CMD ["./app"]
