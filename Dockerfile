FROM golang:latest
ENV PORT=4001
ENV PGHOST=localhost 
ENV PGUSER=user
ENV PGPASSWORD=password
ENV PGDATABASE=dbname
ENV PGPORT=5432
ENV PGPASSWORD_DOCKER=password
ENV PGDATABASE_DOCKER=dbname
WORKDIR /app
COPY go.mod go.sum ./
COPY . .
RUN go mod tidy
RUN go build -o app
RUN go run migrations/migrations.go
CMD ["./app"]
