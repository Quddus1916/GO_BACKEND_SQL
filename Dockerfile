 FROM golang:1.18.1-alpine
#destination
 WORKDIR /dockerapp
#copy mod and sum for dependencies
 COPY go.mod ./
 COPY go.sum ./
#download dependencies
 RUN apk add git
 RUN go mod download
 
 

#copy all other files 
 COPY *.go ./

 RUN go build -o /docker-gs-ping

 CMD ["/docker-gs-ping"]