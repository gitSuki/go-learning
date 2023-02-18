# build stage
FROM golang:1.18.1-alpine3.14 AS builder
WORKDIR /app
COPY . .
RUN cd simple-bank; go build -o main main.go

# run stage
FROM alpine:3.14
WORKDIR /app
COPY --from=builder /app/simple-bank/main .


EXPOSE 8000
CMD [ "/app/main" ]