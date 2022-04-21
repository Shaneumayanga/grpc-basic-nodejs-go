FROM golang:1.17-alpine as builder

RUN apk update && apk add git

WORKDIR /go/src/app
COPY . .
# RUN go mod init github.com/Shaneumayanga/grpc-react
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app .

FROM scratch
WORKDIR /bin
COPY --from=builder /go/bin/app /app

ENTRYPOINT ["/app"]