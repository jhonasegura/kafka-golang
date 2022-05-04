FROM golang:alpine AS build
LABEL maintainer=Team

WORKDIR /app
COPY . .
RUN make build

FROM alpine

WORKDIR /app
COPY --from=build /app/build/bin/golang-with-kafka /app

ENTRYPOINT [ "/app/golang-with-kafka" ]