##
## Build Users
##
FROM golang:latest as build_server_users
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server_users ./servers/users/main.go

##
## Build Articles
##
FROM golang:latest as build_server_articles
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server_articles ./servers/articles/main.go


##
## Deploy Users
##
FROM alpine:latest as server_users
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build_server_users /app/server_users .
EXPOSE 9000
CMD ["./server_users"] 


##
## Deploy Articles
##
FROM alpine:latest as server_articles
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build_server_articles /app/server_articles .
EXPOSE 9000
CMD ["./server_articles"] 