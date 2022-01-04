FROM golang as builder

WORKDIR /Users/adsorbentkarma/Downloads/Drop/Docker1/

COPY . .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

LABEL author="TrestX"

WORKDIR /root/

USER root

COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/Docker1/ .

EXPOSE 6000

CMD ["sh","runservices.sh"]

