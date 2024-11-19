FROM golang:latest

WORKDIR /app/demo
COPY . .

RUN go build eshop_user

EXPOSE 8888
ENTRYPOINT ["./eshop_user"]