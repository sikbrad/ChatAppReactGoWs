#FROM golang:1.11.1-alpine3.8
FROM golang:1.15.6-alpine3.12
RUN mkdir /app
ADD . /app/
WORKDIR /app

# added. git error. fix.
# @reef https://github.com/nodejs/docker-node/issues/586
RUN apk add --no-cache git


RUN go mod download
#RUN go build -o main ./... # it does not builds
RUN go build -o main .
CMD ["/app/main"]