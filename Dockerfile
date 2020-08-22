# FROM golang
# COPY . /go/src/auth
# WORKDIR /go/src/auth
# RUN go get .
# ENTRYPOINT go run main.go
# EXPOSE 3000


# FROM golang:alpine
# RUN mkdir /app 
# ADD . /app/
# WORKDIR /app 
# RUN go build -o main .
# RUN adduser -S -D -H -h /app appuser
# USER appuser
# CMD ["./main"]

FROM golang:1.11.1-alpine3.8 as build-env
# All these steps will be cached
RUN mkdir /app
WORKDIR /app
COPY go.mod . 
# ^ COPY go.mod and go.sum files to the workspace
COPY go.sum .

# Get dependancies - will also be cached if we won't change mod/sum
RUN go mod download
# COPY the source code as the last step
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/app
FROM scratch 
# ^ Second step to build minimal image
COPY --from=build-env /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]