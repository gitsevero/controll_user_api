FROM golang:alpine as builder
WORKDIR /
COPY . .
RUN go build -o main .
FROM scratch
WORKDIR /
COPY --from=builder /main /main
COPY /.env .env
CMD ["/main"]
