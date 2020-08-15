#MultiStage DockerFile
#Developer Sunish Kannembath

FROM golang:alpine AS builder

WORKDIR /go/src/app

COPY /RestAPI/api.go .

RUN go build -o webservice .

FROM alpine

WORKDIR /app

COPY --from=builder /go/src/app /app/

CMD ["./webservice"]