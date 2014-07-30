FROM google/golang

WORKDIR /gopath/src/github.com/rigger-dot-io/rigger-host
ADD     . /gopath/src/github.com/rigger-dot-io/rigger-host

RUN     go get ./...
RUN     make all

CMD []
ENTRYPOINT ["/gopath/bin/rigger"]