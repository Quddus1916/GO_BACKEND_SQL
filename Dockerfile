 FROM golang:1.18
#destination
 WORKDIR /app
#copy mod and sum for dependencies
 COPY go.mod ./
 COPY go.sum ./
#download dependencies
 
 RUN go mod download
#copy all other files 
 COPY . .




 RUN go build -o /docker-echo-app

 CMD ["/docker-echo-app"]

