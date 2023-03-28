FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /ipfs-load-gen

EXPOSE 9100

CMD [ "/ipfs-load-gen" ]