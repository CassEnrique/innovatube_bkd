## Release of Test Versions
## v0.1.0-a.1

## Builder Segment
FROM golang:alpine AS build
RUN apk update
RUN apk add --no-cache git upx

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go get -d -v ./...
RUN go build \
    -ldflags="-s -w" \
    -o ./bkd \
    -v ./cmd/*.go
RUN cp cmd/.env.dev .env
RUN upx ./bkd


## Build Deployment
FROM alpine:latest
#LABEL Name=EmbebedEspiralPackage Version=0.1.0-b.3

RUN apk update
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=build /app .
RUN cp cmd/.env.dev .env

ENTRYPOINT ["./bkd"]

CMD ["-h"]

EXPOSE 1991
