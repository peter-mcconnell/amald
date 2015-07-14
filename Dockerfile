FROM golang:1.5

COPY .git/ /src
RUN cd /src && git archive HEAD > archive.tar
RUN tar -xvf /src/archive.tar -C /src/clean/
WORKDIR /src/clean

CMD ["go test ./..."]
