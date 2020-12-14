FROM alpine:3.12.0

COPY cars .
COPY build/*.dart.js dist/js/
COPY views views

RUN mkdir -p /views/_shared

EXPOSE 8081

ENTRYPOINT [ "./cars" ]
