FROM golang:1.15.7-alpine
WORKDIR /go/src/github.com/Instabug/curly
COPY . .
RUN apk --no-cache add make
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o web .

FROM golang:1.15.7-alpine
WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=0 /go/src/github.com/Instabug/curly/web .
CMD ["./web"]