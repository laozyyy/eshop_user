FROM golang:latest

WORKDIR /app/demo
COPY . .

RUN go build eshop_user

EXPOSE 9000
ENTRYPOINT ["./eshop_user"]