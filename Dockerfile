FROM golang:1.16 as build

ENV GOPROXY=https://goproxy.io

ADD . /car

WORKDIR /car

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api_server

FROM alpine:3.7


ENV GIN_MODE="release"
ENV PORT=3000


RUN mkdir -p /www/conf

WORKDIR /www

COPY --from=build /car/api_server /usr/bin/api_server
ADD ./conf /www/conf

RUN chmod +x /usr/bin/api_server

ENTRYPOINT ["api_server"]