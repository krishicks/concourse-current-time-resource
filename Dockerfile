FROM golang:alpine as builder
COPY . /go/src/github.com/krishicks/concourse-current-time-resource
ENV CGO_ENABLED 0
RUN go build -o /assets/out github.com/krishicks/concourse-current-time-resource/cmd/out
RUN go build -o /assets/in github.com/krishicks/concourse-current-time-resource/cmd/in
RUN go build -o /assets/check github.com/krishicks/concourse-current-time-resource/cmd/check

FROM alpine:edge
COPY --from=builder /assets /opt/resource
