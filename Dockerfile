FROM golang:alpine
WORKDIR /usr/src/app
COPY . .

RUN go build -o main ./app
EXPOSE 8080
CMD ["./main"]
# # Run application
# CMD ["go", "run", "app/main.go"]