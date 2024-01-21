FROM golang:1.21.1

WORKDIR /app

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY . .

RUN go build -o main ./cmd/web

EXPOSE 8080

CMD [ "./main" ]