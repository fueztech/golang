FROM golang:1.16-alpine

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/fueztech/golang/13_webapi/main
RUN cd /build && git clone https://github.com/fueztech/golang/13_webapi/gowebapi.git

RUN cd /build/GoWebAPI/main && go build

EXPOSE 8040

ENTRYPOINT [ "/build/GoWebAPI/main/main"]
