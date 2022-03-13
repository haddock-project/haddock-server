# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:alpine AS build

WORKDIR /app

COPY . .
RUN go mod download

# needed to run cgo packages
RUN apk add build-base

RUN go build -o /Haddock

##
## Deploy
##
FROM alpine

WORKDIR /

COPY --from=build /Haddock /Haddock
RUN chmod +x /Haddock

EXPOSE 8080

ENTRYPOINT ["/Haddock"]