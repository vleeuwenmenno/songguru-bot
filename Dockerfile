FROM golang:1.20.8-alpine
RUN apk add --no-cache gcc g++ git openssh-client

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /app/songshizz_bot

EXPOSE 8080
CMD ["/app/songshizz_bot"]