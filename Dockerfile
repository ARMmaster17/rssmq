FROM golang:1.17-alpine
RUN mkdir /src
ADD ./* /src
WORKDIR /src
RUN go build ./cmd/main.go
CMD /src/cmd/main