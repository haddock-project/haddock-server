# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:alpine AS build

WORKDIR /app

COPY ["go.sum", "go.mod", "./"]
RUN go mod download

COPY . .

RUN go build -o /kontainerized

##
## Deploy
##
FROM alpine

WORKDIR /

COPY --from=build /kontainerized /kontainerized
RUN chmod +x /kontainerized

EXPOSE 8080

ENTRYPOINT ["/kontainerized"]