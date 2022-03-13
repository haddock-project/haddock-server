# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:alpine AS build

WORKDIR /app

COPY ["go.sum", "go.mod", "./"]
RUN go mod download

COPY . .

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