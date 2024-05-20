FROM golang:1.22

# set working directory inside the container
WORKDIR /go/src/app

# copy go.mod and go.sum so we can download dependencies first
COPY go.mod go.sum ./

# download all the dependencies / packages
RUN go mod download

# copy sources to the working directory
COPY . .

# build the go project 
RUN go build -o /mutasi-service ./cmd/web/*.go

# copy .env file
# COPY config.env /build

# listen to this port
EXPOSE 20027

# RUN the app
CMD ["/mutasi-service"]