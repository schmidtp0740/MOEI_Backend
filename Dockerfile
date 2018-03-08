FROM golang

WORKDIR /go/src/app
COPY . .
RUN go get github.com/gorilla/mux &&  go build -o app
CMD [ "./app" ]