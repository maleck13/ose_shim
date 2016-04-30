FROM golang:1.6

ADD . /go/src/github.com/maleck13/ose_shim

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN cd /go/src/github.com/maleck13/ose_shim && go get . &&  go install .

RUN curl -L https://get.docker.com/builds/Linux/x86_64/docker-1.8.2 > /usr/bin/docker && chmod +x /usr/bin/docker

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/ose_shim serve --config=/go/src/github.com/maleck13/ose_shim/config/config.json

# Document that the service listens on port 8080.
EXPOSE 3000