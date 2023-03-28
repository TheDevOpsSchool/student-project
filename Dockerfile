FROM golang

WORKDIR /microapp

COPY /authentication.go .

RUN go mod init https://github.com/TheDevOpsSchool/student-project.git


RUN go mod tidy


RUN go mod download && go mod verify

RUN go build -o authentication.go

EXPOSE 8080

CMD ["./authentication.go"]
