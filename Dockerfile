FROM golang:1.12.4 as build

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

FROM gcr.io/distroless/base
COPY --from=build /go/src/app/data /data
COPY --from=build /go/bin/app /
CMD ["/app"]