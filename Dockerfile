FROM golang:1.17.6-alpine AS builder
WORKDIR /app/
COPY go.mod go.sum ./  
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

FROM alpine:latest  
WORKDIR /root/
COPY --from=builder /app/static/ ./static/
COPY --from=builder /app/main ./
EXPOSE 8080
CMD ["./main"]  
