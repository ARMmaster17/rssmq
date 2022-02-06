FROM golang:1.17-alpine as go-build-stage
RUN mkdir /src
COPY ./ /src/
WORKDIR /src
RUN go build ./main.go -o rssmq

FROM alpine:3.15.0 as final-stage
COPY --from=go-build-stage /src/rssmq /rssmq
CMD rssmq
