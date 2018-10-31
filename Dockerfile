FROM golang:1.11
WORKDIR /go/src/app
ADD vendor vendor
ADD dao.go Gopkg.lock Gopkg.toml hack.go heartRate.go helper.go ins.go main.go people.go request.go rx.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
COPY --from=0 /go/src/app/app .
EXPOSE 80
CMD ["./app"]
