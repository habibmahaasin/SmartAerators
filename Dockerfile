FROM golang:alpine
WORKDIR /usr/src/app
COPY . .

RUN go mod tidy
RUN go build .app/main.go

EXPOSE 8080

ENTRYPOINT [ "./main" ]
# CMD ["./main"]
# CMD ["go", "run", "app/main.go"]