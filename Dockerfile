FROM golang:1.10.1

# Build contributor-counter
ADD . /go/src/github.com/fossas/contributor-counter
RUN [ "go", "install", "github.com/fossas/contributor-counter/..." ]

# Remove SSL root CAs to test invalid certificate override
RUN [ "rm", "-rf", "/etc/ssl/certs" ]
RUN [ "mkdir", "-p", "/etc/ssl/certs" ]

CMD [ "/bin/bash" ]
