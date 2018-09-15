FROM golang:1.10
USER 1001
WORKDIR $GOPATH/src/HZ_proj/Backend
COPY . /go/src/HZ_proj/Backend

RUN go get -u github.com/go-sql-driver/mysql
RUN go install 
EXPOSE 7777
ENTRYPOINT ["/go/bin/HZ_proj"]
