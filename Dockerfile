FROM golang:1.13.5 as build

WORKDIR /go/src/github.com/JulianSauer/FritzPowerMonitor
ADD . /go/src/github.com/JulianSauer/FritzPowerMonitor

RUN go get -d -v ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:3.12 as certificates

RUN apk add -U --no-cache ca-certificates

FROM scratch
WORKDIR /bin
COPY --from=build /go/src/github.com/JulianSauer/FritzPowerMonitor/main .
COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

ENTRYPOINT ["./main"]
