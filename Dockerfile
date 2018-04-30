# Use oldest supported Ubuntu
FROM ubuntu:14.04

# Install build tools
RUN [ "apt-get", "update" ]
RUN [ "apt-get", "install", "-y", "wget", "git=1:1.9.1-1" ]

# Install latest Go
WORKDIR /tmp/go1.10.1
RUN [ "wget", "https://dl.google.com/go/go1.10.1.linux-amd64.tar.gz" ]
RUN [ "tar", "-xf", "go1.10.1.linux-amd64.tar.gz" ]
RUN [ "mv", "go", "/usr/local" ]
ENV GOROOT=/usr/local/go GOPATH=$HOME/go
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

# Build newer version of `git` from source to test shallow cloning
WORKDIR /tmp/git/2.17.0
RUN [ "apt-get", "install", "-y", "dh-autoreconf", "libcurl4-gnutls-dev", "libexpat1-dev", "gettext", "libz-dev", "libssl-dev" ]
RUN [ "wget", "https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.17.0.tar.gz" ]
RUN [ "mkdir", "-p", "/tmp/git/2.17.0/src" ]
RUN [ "tar", "-xf", "git-2.17.0.tar.gz", "--directory", "src", "--strip-components=1"]
WORKDIR /tmp/git/2.17.0/src
RUN [ "make", "configure" ]
RUN [ "./configure", "--prefix=/tmp/git/2.17.0" ]
RUN [ "make", "all" ]
RUN [ "make", "install" ]

# Remove SSL root CAs to test invalid certificate override
RUN [ "apt-get", "remove", "-y", "ca-certificates" ]
RUN [ "rm", "-rf", "/etc/ssl/certs" ]
RUN [ "mkdir", "-p", "/etc/ssl/certs" ]

# Build contributor-counter
ADD . /go/src/github.com/fossas/contributor-counter
RUN [ "go", "install", "github.com/fossas/contributor-counter/..." ]

WORKDIR /root
CMD [ "/bin/bash" ]
