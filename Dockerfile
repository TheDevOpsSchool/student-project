FROM golang

WORKDIR /uploadapp

COPY /file.go .

RUN go mod init https://github.com/TheDevOpsSchool/student-project.git


RUN go mod tidy


RUN go mod download && go mod verify

RUN go build -o file.go

EXPOSE 8080

CMD ["./file.go"]
