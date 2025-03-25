FROM golang:alpine AS builder

WORKDIR /app

COPY ./go.mod ./go.mod
COPY ./go.sum ./go.sum

COPY ./schema ./schema
COPY ./web ./web
COPY ./go ./go

RUN go build -o /server ./go

FROM alpine

WORKDIR /app

COPY --from=builder /server ./server

ARG PORT=8080
EXPOSE ${PORT}

CMD ["./server"]
