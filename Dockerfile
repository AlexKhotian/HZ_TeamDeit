FROM golang:1.10
USER 1001
WORKDIR .
COPY . .

RUN go install 
EXPOSE 7777
ENTRYPOINT ["/go/bin/HZ_proj"]
