FROM golang:latest AS builder

ARG PORT=10500
ENV PORT=$PORT

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /iam-api && echo "Finished build"

EXPOSE $PORT

# Run the compiled binary
CMD ["/iam-api"]
