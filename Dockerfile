FROM alpine:latest

COPY cars .
COPY conf conf
COPY views views
COPY static static

##Download the latest templates
RUN apk add --update curl && rm -rf /var/cache/apk/*
RUN apk --no-cache add jq
RUN for k in $(curl -XGET localhost:8093/v1/asset/html | jq ".Data | .[]"); do curl -O views/_shared/$k localhost:8093/v1/asset/html/$k; done

EXPOSE 8081

CMD [ "./cars"]