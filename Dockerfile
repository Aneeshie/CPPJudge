FROM golang:1.22

WORKDIR /app

RUN go install github.com/air-verse/air@latest

#cache
COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

CMD ["air"]
