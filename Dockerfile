FROM golang

WORKDIR /$GOPATH/src/github.com/schmidtp0740/moei_backend
RUN go get github.com/gorilla/mux && \
  go get github.com/rs/cors
COPY . .
RUN go build -o app
CMD [ "./app" ]
