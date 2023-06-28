# Build stage
FROM golang:1.19.9-alpine AS build
WORKDIR /app
COPY . .
RUN go build -o myapp

# Production stage
FROM alpine:3.14
WORKDIR /app
COPY --from=build /app/myapp .
CMD ["./myapp"]
