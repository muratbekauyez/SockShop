FROM golang:1.16.2-alpine as build

WORKDIR /Users

COPY . .

RUN mkdir /app
WORKDIR /app

FROM alpine
WORKDIR /app
COPY --from=build /Users/api .

CMD ["./Users"]