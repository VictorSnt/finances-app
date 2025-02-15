# Usar uma imagem oficial do Go
FROM golang:1.23.5

# Definir o diretório de trabalho dentro do container
WORKDIR /app

# Copiar os arquivos do projeto para dentro do container
COPY . /app

# Instalar o swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Construir o binário do aplicativo e gerar a documentação Swagger
RUN go build -o main main.go && swag init

# Expor a porta 80
EXPOSE 80

# Comando para rodar o aplicativo
CMD ["./main"]