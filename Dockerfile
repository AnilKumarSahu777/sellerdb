FROM golang:1.12-alpine

WORKDIR /

COPY sellerdb /
CMD ["/sellerdb"]
