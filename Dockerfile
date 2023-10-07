FROM golang:1.20.8-alpine
RUN apk add --no-cache gcc g++ git openssh-client

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN rm -rf configs/
RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux go build -ldflags="-w -s" -o /app/songshizz_bot

HEALTHCHECK --interval=10s --timeout=5s --start-period=30s \ 
  CMD wget --no-verbose --tries=1 http://localhost:8080/changelogs -O /dev/null || exit 1  

EXPOSE 8080
CMD ["/app/songshizz_bot"]