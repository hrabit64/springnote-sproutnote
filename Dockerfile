FROM golang:1.23
WORKDIR /app
COPY . .
RUN mkdir ./logs
RUN mkdir ./backup

RUN go build -o /bin/sproutnote ./cmd/cil/main.go
RUN go build -o /bin/sproutnote-scheduler ./cmd/scheduler/main.go

RUN chmod +x /bin/sproutnote
RUN chmod +x /bin/sproutnote-scheduler

RUN apt-get -y update && apt-get install -y mariadb-client bash

CMD ["/bin/sproutnote-scheduler"]