FROM golang:1.18

RUN mkdir /task
COPY . /task

ENV GODEBUG madvdontneed=1

WORKDIR /task/cmd/admin

RUN cd /task/cmd/admin
RUN go mod verify && go build -o admin
CMD ["./admin"]