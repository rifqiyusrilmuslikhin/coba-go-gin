FROM golang:latest

ENV PORT=4001
ENV PGHOST=localhost 
ENV PGUSER=developer
ENV PGPASSWORD=supersecretpassword
ENV PGDATABASE=coba_golang
ENV PGPORT=5432

ENV PGPASSWORD_DOCKER=sembarang
ENV PGDATABASE_DOCKER=coba_golang

WORKDIR /app
COPY go.mod go.sum ./

COPY . .

RUN go mod tidy

RUN go build -o app

RUN go run migrations/migrations.go

CMD ["./app"]
