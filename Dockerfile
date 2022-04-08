FROM golang:1.17

LABEL maintainer 'Brenddon Anjos <brenddon.dev@gmail.com>' 

WORKDIR /go/src

#Includes air to hot reload
RUN go get github.com/cosmtrek/air
