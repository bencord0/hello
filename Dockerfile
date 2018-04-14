FROM golang as builder
ADD main.go ./
RUN go build -v -o /hello main.go

EXPOSE 80
CMD ["/hello"]
