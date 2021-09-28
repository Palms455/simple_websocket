FROM golang:latest as builder
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/chat/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 5555
CMD ["/app/main"]