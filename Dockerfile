# Usar la imagen oficial de Go
FROM golang:1.24.1-alpine

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

COPY go.mod .
RUN go mod download

# Copiar los archivos de la aplicación
COPY . .

RUN go build -o out/hello .
# Exponer el puerto 8080
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["out/hello"]
