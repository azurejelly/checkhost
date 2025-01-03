FROM golang:1.23.2 AS build

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 go build -o /build/checkhost

FROM gcr.io/distroless/static-debian12

WORKDIR /app

COPY --from=build /build/checkhost .

ENTRYPOINT [ "./checkhost" ]