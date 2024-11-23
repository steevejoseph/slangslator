# Dev
# FROM golang:1.23-alpine
# WORKDIR /app
# COPY . .
# RUN go CGO_ENABLED=0 build -o slangslator
# CMD ["./myapp"]

# PROD
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o myapp

FROM alpine:latest
COPY --from=builder /app/myapp /
ENTRYPOINT ["/myapp"]
