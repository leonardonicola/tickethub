# Build the application from source
FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Run the tests in the container
FROM build-stage AS test
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM alpine:latest AS release

WORKDIR /root/

COPY --from=build-stage /app/.env .
COPY --from=build-stage /app/main .

EXPOSE 3000

CMD ["./main"]