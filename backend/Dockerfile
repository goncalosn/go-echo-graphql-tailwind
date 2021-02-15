FROM golang:1.14
WORKDIR /go/src/app
COPY . .
RUN go run ./cmd/main.go
EXPOSE 1323


# 1. fazes o FROM para ir buscar a imagem de golang
# 2. depois copias os ficheiros locais para esta imagem
# 3. agora com a copia dos ficheiros, já tens o go.mod e aqui estao as dependencies, logo sacas as dependencies
# 4. o passo seguinte é fazer o go install/ go build para criar o binario da app
