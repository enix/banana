FROM golang:1.11 as prod

WORKDIR /go/src/enix.io/banana

COPY . .

ADD https://raw.githubusercontent.com/golang/dep/master/install.sh /usr/local/bin/install-dep

RUN chmod +x /usr/local/bin/install-dep

RUN /usr/local/bin/install-dep

RUN dep init

RUN go build -v -o banana ./src

EXPOSE 80

CMD ["/go/src/enix.io/banana/banana"]

# ---------------------------

FROM golang:1.11

COPY --from=prod /go /go

COPY --from=prod /usr/local/go /usr/local/go

RUN go get github.com/canthefason/go-watcher

RUN go install github.com/canthefason/go-watcher/cmd/watcher

CMD ["bash", "-c", "cd /go/src/enix.io/banana/src && watcher"]
