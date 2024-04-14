FROM golang:1.22-alpine

WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN apk add --no-cache make

RUN make dependencies setup generate-mocks

CMD ["make" ,"unit-tests"]

