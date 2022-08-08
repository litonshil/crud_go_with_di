FROM golang:1.18-alpine
ENV GO111MODULE=on

RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add git

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

# COPY *.go ./

RUN go build -o /crud_go .
EXPOSE 8080
CMD ["/crud_go"]

# command:

#docker build -t crud_go_with_di