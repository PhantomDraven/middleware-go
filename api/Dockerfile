FROM golang:latest

WORKDIR /app

RUN go install github.com/air-verse/air@latest
COPY .air.toml ./
COPY . .
RUN go mod download

CMD ["air", "-c", ".air.toml"]
