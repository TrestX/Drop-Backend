# FROM golang as builder

# WORKDIR /Users/adsorbentkarma/Downloads/Drop/DropAddress
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o address .
# WORKDIR /Users/adsorbentkarma/Downloads/Drop/DropShop
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shop .
# WORKDIR /Users/adsorbentkarma/Downloads/Drop/DropUserAccount
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o account .
# WORKDIR /Users/adsorbentkarma/Downloads/Drop/DropWallet
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o wallet .

# FROM alpine:latest  
# RUN apk --no-cache add ca-certificates

# LABEL author="TrestX"

# WORKDIR /root/
# USER root 

# COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/DropUserAccount/account .
# COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/DropWallet/wallet .
# COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/DropAddress/address .
# COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/DropShop/shop .
# COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/DropUserAccount/conf .
# COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/DropUserAccount/runservices.sh .
# RUN chmod 755 /root/runservices.sh
# CMD /root/runservices.sh

# EXPOSE 6009 6010 6025 6029

#Docker 1
# FROM golang as builder

# WORKDIR /Users/adsorbentkarma/Downloads/Drop/Docker1/

# COPY . .

# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# LABEL author="TrestX"

# WORKDIR /root/

# USER root

# COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/Docker1/ .

# EXPOSE 6009 6010 6025 6029

# CMD ["sh","runservices.sh"]


#Docker 2

# FROM golang as builder

# WORKDIR /Users/adsorbentkarma/Downloads/Drop/Docker2/

# COPY . .

# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# LABEL author="TrestX"

# WORKDIR /root/

# USER root

# COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/Docker2/ .

# EXPOSE 6020 6031 6021

# CMD ["sh","runservices.sh"]


#Docker 3

# FROM golang as builder

# WORKDIR /Users/adsorbentkarma/Downloads/Drop/Docker3/

# COPY . .

# FROM alpine:latest

# RUN apk --no-cache add ca-certificates

# LABEL author="TrestX"

# WORKDIR /root/

# USER root

# COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/Docker3/ .

# EXPOSE 6012 6018 6011 6033

# CMD ["sh","runservices.sh"]

#Docker 4

FROM golang as builder

WORKDIR /Users/adsorbentkarma/Downloads/Drop/Docker6/

COPY . .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

LABEL author="TrestX"

WORKDIR /root/

USER root

COPY --from=builder /Users/adsorbentkarma/Downloads/Drop/Docker6/ .

EXPOSE 6027

CMD ["sh","runservices.sh"]
