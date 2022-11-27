FROM golang:1.18-bullseye

RUN go install github.com/gin-gonic/gin

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

ENV APP_HOME /go/src/go_todo
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"
EXPOSE 8080
CMD ["gin", "run"]
