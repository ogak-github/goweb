# Go Image version same as my local machine
FROM golang:1.23.4-alpine AS builder
# App Folder
WORKDIR /app
# Copy file project
COPY . .
# Download dependency
RUN go mod download
# Build go project
RUN go build -o /bin/app


# Run Stage
FROM alpine:latest
COPY --from=builder /bin/app /bin/app
EXPOSE 8080
CMD ["/bin/app"]
