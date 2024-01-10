FROM golang:1.22rc1-bookworm
WORKDIR /app
COPY . .
RUN go install
EXPOSE 8080
ENTRYPOINT go run main.go