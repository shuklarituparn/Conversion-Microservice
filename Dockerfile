
FROM golang:1.22-alpine

RUN apk update && \
    apk add --no-cache gcc libc-dev build-base ffmpeg

WORKDIR /app


RUN apk update && apk add --no-cache gcc libc-dev build-base


COPY . .


RUN go mod download

WORKDIR /app/cmd/video-converter


RUN CGO_ENABLED=1 GOOS=linux go build -tags musl -o main


EXPOSE 8085


CMD ["./main"]