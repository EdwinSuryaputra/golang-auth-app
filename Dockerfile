FROM golang:latest AS builder

ARG PORT=10500
ENV PORT=$PORT

WORKDIR /app

COPY . .

RUN apt-get update && apt-get install -y netcat-openbsd
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /golang-auth-app && echo "Finished build"

RUN chmod +x entrypoint.sh

EXPOSE $PORT

CMD ["sh", "entrypoint.sh"]
