FROM golang:alpine
WORKDIR /usr/src/app
COPY . .

RUN go mod tidy
RUN go build .app/main.go

EXPOSE 6000

#CMD ["make", "start"]
ENTRYPOINT [ "./main" ]
# # Run application
# CMD ["go", "run", "app/main.go"]