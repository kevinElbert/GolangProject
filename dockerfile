# Gunakan base image Golang
FROM golang:1.22-alpine

# Set working directory
WORKDIR /app

# Copy go.mod dan go.sum untuk dependency management
COPY go.mod go.sum ./

# Download module dependencies
RUN go mod download

# Copy seluruh kode sumber ke dalam container
COPY . .

# Compile aplikasi ke binary
RUN go build -o main .

# Ekspos port yang digunakan aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["/app/main"]
