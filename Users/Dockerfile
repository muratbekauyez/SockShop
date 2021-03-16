FROM golang:1.16.2-alpine as build

WORKDIR /go/src/app

COPY . .

RUN go run -o main ./main.go

FROM alpine
WORKDIR /app
COPY --from=build /go/src/app/ /app/

CMD ["/app/main"]