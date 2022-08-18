FROM golang:1.18

RUN mkdir /task
COPY . /task

ENV GODEBUG madvdontneed=1

WORKDIR /task/cmd/executor

RUN cd /task/cmd/executor
RUN go env -w GO111MODULE=on && go env -w GOPROXY="https://goproxy.cn,direct"

RUN go mod tidy && go build -o executor
CMD ["./executor"]