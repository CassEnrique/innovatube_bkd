# InnovaTube - API

Esta es una aplicación desarrollada en [Golang](https://go.dev/) con los frameworks [Gorm](https://gorm.io/) como
orm para Golang, con soporte para diversos [gestores de bases de datos](https://www.hostinger.mx/tutoriales/sgbd),
y [Fiber](https://docs.gofiber.io/) como web framework.

## Requisitos

[Golang](https://go.dev/dl/)
[Docker](https://www.docker.com/products/docker-desktop/)

## Clonar

Para comenzar, haz un clon del repositorio [InnovaTube-Bkd](https://github.com/CassEnrique/innovatube_bkd) en su
rama main:

```bash
git clone -b main git@github.com:CassEnrique/innovatube_bkd.git innovatube_bkd
cd innovatube_bkd
```

## Colaborar

Si vas a desarrollar, y aún no existe una rma `Dev`, contribuye creando esta con los siguientes comandos y colabora
para dar más valor al proyecto:

```bash
git clone -b main git@github.com:CassEnrique/innovatube_bkd.git innovatube_bkd
git checkout -b development
# develop
git add .
git pull
git push origin development
```

## Ejecución

Para ejecutar la aplicación en desarrollo, primero copio el archivo `.env.dev` del directorio `./cmd`:

```bash
# Linux
cp ./cmd/.env.dev ./.env
# Windows
copy .\cmd\.env.dev .\.env
New-Item -Path .\cmd\.env.dev -Destination .\.env
```

Descarga las dependencias:
```bash
go mod tidy
```

Recomendado instalar [air](https://github.com/air-verse/air), para ello el siguiente comando:
```bash
go install github.com/air-verse/air@latest
```

## Instanciar imagen en docker
```bash
docker run --name postgres-dev -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -e POSTGRES_DB=dm_innovatube -p 2345:5432 -d postgres
```

## Ejecutamos air
```bash
air init
# Modificas el archivo air en la seccion [Build] - bin y cmd si estas en Windows
[build]
  bin = ".\bkd.exe"
  cmd = "go build -o .\bkd.exe .\cmd\."
# Si estas en Linux, ejecuta
air
```

## Ejecutar sin air
```bash
# Windows
go build -o .\bkd.exe .\cmd\.
.\bkd.exe
# Linux
go build -o ./bkd ./cmd/*.go
./bkd
```

La aplicación esta corriendo por defecto en el puerto `1991`
