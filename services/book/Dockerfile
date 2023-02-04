FROM golang:1.19-alpine

WORKDIR /go/src/
COPY ./services/book .

RUN apk upgrade --update && apk --no-cache add git

COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -o /usr/bin/book ./services/book
RUN go get -u github.com/cosmtrek/air && go build -o /go/bin/air github.com/cosmtrek/air

CMD ["air", "-c", ".air.toml"]
