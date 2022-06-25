FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download github.com/sagikazarmark/crypt
RUN go mod download

COPY . ./

RUN go build -o /application

CMD [ "/application" ]

