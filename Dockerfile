FROM golang:1.10 AS build
WORKDIR /go/src/github.com/sachinmaharana/go-app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .
CMD ["./go-app"]

# Prod

FROM scratch as prod
COPY --from=build /go/src/github.com/sachinmaharana/go-app .
CMD ["./go-app"]

