FROM golang:1.17
WORKDIR /go/src

RUN apt-get update && \
    apt-get install nodejs npm -y && \
    apt-get install sqlite3 -y && \
    npm install -g nodemon

COPY . .

ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=1
ENV GO111MODULE=on

CMD ["tail", "-f", "/dev/null"]
CMD nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/api/main.go