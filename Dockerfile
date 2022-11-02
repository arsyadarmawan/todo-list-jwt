FROM golang:alpine

RUN apk update && apk add --no-cache gcc libc-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

ADD . .

RUN go build -o /bin/moonlay ./main.go

WORKDIR /

COPY .env .

EXPOSE 3000

CMD /bin/moonlay