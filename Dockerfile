FROM golang:1.17-alpine
RUN mkdir /src
COPY ./ /src/
WORKDIR /src
RUN ls -al
RUN go build ./cmd/main.go
RUN ls -al
CMD /src/main