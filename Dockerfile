FROM golang:1.13 as build_base

WORKDIR /box

COPY go.mod .
COPY go.sum .

RUN go mod download

FROM build_base as builder

COPY main.go .
COPY handles ./handles

RUN CGO_ENABLED="0" go build

FROM google/dart AS pyltjie
ENV PATH="$PATH:/root/.pub-cache/bin"

WORKDIR /arrow
RUN pub global activate webdev

COPY build.yaml build.yaml
COPY pubspec.yaml pubspec.yaml
RUN pub get

COPY web ./web
COPY lib ./lib
RUN webdev build

FROM alpine:3.12.0

COPY --from=builder /box/cars .
COPY --from=pyltjie /arrow/build/*.dart.js dist/js/
COPY views views

RUN mkdir -p /views/_shared

EXPOSE 8081

ENTRYPOINT [ "./cars" ]
