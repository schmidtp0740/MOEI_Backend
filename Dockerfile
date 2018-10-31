FROM golang:1.11
WORKDIR /go/src/app
ADD . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
COPY --from=0 /go/src/app/app .
EXPOSE 80
CMD ["./app"]
