FROM golang:1.8

WORKDIR /go/src/app

RUN go get -v github.com/roundpartner/seq
RUN go install -v github.com/roundpartner/seq

EXPOSE 6060

CMD ["seq"]