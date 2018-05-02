FROM golang:1.7

RUN mkdir /app
COPY ./vendor/manifest /go/src/github.com/microservices-demo/payment/vendor/manifest

RUN go get -u github.com/FiloSottile/gvt
RUN cd /go/src/github.com/microservices-demo/payment/ && gvt restore

COPY . /go/src/github.com/microservices-demo/payment/

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main github.com/microservices-demo/payment/cmd/paymentsvc

CMD ["/app/main", "-port=80"]

#EXPOSE 80
