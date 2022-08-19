FROM golang:1.19-alpine AS BUILD

WORKDIR /app

RUN apk add build-base

COPY go.mod /app
COPY go.sum /app

RUN go mod download

COPY . /app

RUN go build -o api

RUN make migrate

FROM alpine:3.16.2

COPY --from=BUILD /app/api /bin/
COPY --from=BUILD /app/sqlite.db /

EXPOSE 8000

CMD ["api"]