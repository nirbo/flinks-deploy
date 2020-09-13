FROM golang:1.14.9-buster

LABEL maintainer="Nir Ben-Or for Flinks <nirbo2@gmail.com>"

RUN mkdir -p /opt/flinks-crud-rest-api
COPY flinks-crud-rest-api/flinks-crud-rest-api /opt/flinks-crud-rest-api/
WORKDIR /opt/flinks-crud-rest-api/
CMD ["./flinks-crud-rest-api"]
