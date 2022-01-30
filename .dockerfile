FROM golang:1.17-bullseye as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go mod download
RUN go build -o /go/bin/app ./src/main.go

FROM gcr.io/distroless/base-debian11
COPY --from=build /go/bin/app /app

ENV PORT 4000
EXPOSE 4000

CMD ["/app"]