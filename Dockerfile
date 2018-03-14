FROM golang

WORKDIR /go/src/app
COPY . .
RUN go get github.com/gorilla/mux && go get github.com/rs/cors &&  go build -o app
CMD [ "./app" ]
