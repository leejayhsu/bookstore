# GO Repo base repo
FROM golang:1.15-alpine3.12 as builder

RUN apk add git

RUN mkdir /app
ADD . /app
WORKDIR /app

# Download all the dependencies
COPY go.mod go.sum ./

RUN go mod download
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# GO Repo base repo
FROM alpine:latest

RUN apk --no-cache add ca-certificates curl
RUN mkdir /app
WORKDIR /app/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
# COPY wait-for .

EXPOSE 5000


# Run Executable
# CMD ["./wait-for-postgres.sh", "./main"]