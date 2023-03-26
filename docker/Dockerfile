FROM golang:1.20.2-alpine

# Install psql
RUN apk update && apk add postgresql-client

WORKDIR /app

COPY app/go.mod ./
COPY app ./

# Add these lines to get the required packages
# RUN go get github.com/gin-gonic/gin
# RUN go get github.com/lib/pq
RUN go install github.com/cosmtrek/air@latest

# Copy go.mod and go.sum
COPY app/go.mod app/go.sum ./

# Download all dependencies
RUN go mod download
RUN go mod tidy

# Copy the wait-for-db.sh script
COPY docker/wait-for-db.sh /wait-for-db.sh
RUN chmod +x /wait-for-db.sh

ENTRYPOINT ["/wait-for-db.sh", "db", "air"]
