FROM golang:1.9.2-alpine3.6 AS build-stage
WORKDIR /go/src/github.com/adrientoub/ical-tvshows
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags "-s -w" -o ical

FROM scratch
WORKDIR /root/
COPY --from=build-stage /go/src/github.com/adrientoub/ical-tvshows/ical .
CMD [ "./ical" ]
