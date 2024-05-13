FROM golang:alpine as builder
WORKDIR /
COPY /app .
RUN go build -o main .

FROM scratch
WORKDIR /
COPY --from=builder /main /main
COPY /app/.env /.env
CMD ["./main"]
