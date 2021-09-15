FROM golang:1.17-alpine
RUN mkdir /src
ADD ./* /src