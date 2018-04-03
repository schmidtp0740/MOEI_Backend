FROM golang

WORKDIR /$GOPATH/src/github.com/schmidtp0740/moei_backend
COPY . .
RUN go get github.com/gorilla/mux &&  \
	go get github.com/rs/cors && \
	go get gopkg.in/mgo.v2
RUN go build -o app
CMD [ "./app" ]
