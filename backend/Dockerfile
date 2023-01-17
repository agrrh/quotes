FROM golang:1.19-alpine as builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o ./app
RUN chmod +x ./app

# ---

FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app/app ./

CMD ["./app"]
