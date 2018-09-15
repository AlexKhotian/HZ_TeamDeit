FROM golang:1.10
USER 1001
WORKDIR $GOPATH/src/HZ_proj
COPY . /go/src/HZ_proj

RUN go get -u github.com/go-sql-driver/mysql
RUN go install 
EXPOSE 7777
ENTRYPOINT ["/go/bin/HZ_proj"]
