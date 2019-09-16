from golang:1.5.3

RUN apt-get update && \
    DEBIAN_FRONTEND=noninteractive apt-get install -y \
    git \
    ssh

RUN mkdir /root/go && chown root: /root/go
ADD garyBusey.tar.gz /root
RUN chown -R root: /root/garyBusey

RUN echo "export GOPATH=\$HOME/GO" >> /root/.bashrc && echo "export PATH=\$PATH\$GOPATH/bin" >> /root/.bashrc
ENV GOPATH=/root/go

WORKDIR /root/garyBusey
RUN go get github.com/gorilla/mux
RUN go get github.com/gorilla/securecookie
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/boltdb/bolt
RUN go build

CMD /root/garyBusey/garyBusey
