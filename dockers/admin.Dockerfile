FROM golang:1.18

RUN mkdir /admin
COPY . /admin
WORKDIR /task/cmd/admin

RUN cd /tnt/cmd/admin
RUN go mod verify && go build -o admin
CMD ["./admin"]