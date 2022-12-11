FROM golang:1.19.3-buster

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o readbook ./main.go

CMD [ "./readbook" ]