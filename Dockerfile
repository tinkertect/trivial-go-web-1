FROM golang:1.25.1 as build

WORKDIR /src

COPY . /src

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/trivial-go-web-1 ./*.go

FROM scratch

COPY --from=build /bin/trivial-go-web-1 /trivial-go-web-1

CMD ["/trivial-go-web-1"]