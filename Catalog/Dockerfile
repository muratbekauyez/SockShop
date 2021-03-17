FROM golang:1.16.2-alpine as build

WORKDIR /Catalog

COPY . .

RUN go mod download

RUN mkdir /app
WORKDIR /app

FROM alpine
WORKDIR /app
COPY --from=build /Catalog/api .

CMD ["./Catalog"]