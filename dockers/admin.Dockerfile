FROM golang:1.18

RUN mkdir /task
COPY . /task

ENV GODEBUG madvdontneed=1

WORKDIR /task/cmd/admin

RUN cd /task/cmd/admin
RUN go env -w GO111MODULE=on && go env -w GOPROXY="https://goproxy.cn,direct"

RUN go mod tidy && go mod verify && go build -o admin
CMD ["./admin"]