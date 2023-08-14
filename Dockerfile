# Version de golang
FROM golang:1.20

# Directorio de trabajo
WORKDIR /src/app

# Dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar codigo fuente
COPY . ./

RUN go build ./cmd/api/main.go

# Puerto
EXPOSE 90

#
CMD ["./src/app"]


