FROM golang:1.20.4-alpine3.17
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server
EXPOSE 8080
CMD ["/server"]