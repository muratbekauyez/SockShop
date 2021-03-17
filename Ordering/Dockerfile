FROM golang:1.16.2-alpine as build

WORKDIR /Ordering

COPY . .

RUN go mod download

RUN mkdir /app
WORKDIR /app

FROM alpine
WORKDIR /app
COPY --from=build /Ordering/api .

CMD ["./Ordering"]