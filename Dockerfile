FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

WORKDIR /app
ENV PATH="/go/bin:${PATH}"

COPY . .

RUN go mod download

RUN GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

EXPOSE 8080

CMD ["./main"]
