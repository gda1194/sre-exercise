# Version de golang
FROM golang:1.20

# Directorio de trabajo
WORKDIR /app

# Dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar codigo fuente
COPY . ./

# Crear ejecutable
RUN go build ./cmd/api/main.go

# Puerto
EXPOSE 90

# Ejectutar
CMD ["./main"]


