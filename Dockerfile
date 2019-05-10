FROM golang:1.11 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY controllers ./controllers
COPY core ./core
COPY routers ./routers

RUN CGO_ENABLED="0" go build

FROM google/dart AS pyltjie

WORKDIR /arrow
COPY static/dart ./assets/dart

RUN mkdir -p assets/js
COPY compiledart.sh .
RUN sh ./compiledart.sh

FROM alpine:latest

COPY --from=builder /box/cars .
COPY --from=pyltjie /arrow/assets/js dist/js
COPY conf conf
COPY views views

RUN mkdir -p /views/_shared

EXPOSE 8081

ENTRYPOINT [ "./cars" ]
