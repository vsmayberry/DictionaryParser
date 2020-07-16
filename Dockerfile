FROM golang:1.14

WORKDIR /go/src/DictionaryParser
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["DictionaryParser"]
EXPOSE 8080
