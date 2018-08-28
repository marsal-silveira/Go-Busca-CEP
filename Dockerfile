# Base this docker container off the official golang docker image. Docker containers inherit everything from their base.
FROM golang

# Allows app_env to be set during build
# ARG app_env
ENV PORT 8081

# Copy the example-app directory (where the Dockerfile lives) into the container and make it the working directory.
COPY ./app /go/src/cep-provider/app
WORKDIR /go/src/cep-provider/app

# Download and install any required third party dependencies into the container.
RUN go get ./
RUN go build

# Now tell Docker what command to run when the container starts
CMD app;

# Set the PORT environment variable inside the container and expose it to the host so we can access our application
EXPOSE $PORT

# Examples to build and run
# docker build ./ -t cep-provider
# docker run --rm -d -p8081:8080 cep-provider
# docker-compose up -d
# docker-compose start










# FROM scratch

# # Allows app_env to be set during build
# ENV PORT 8080

# # Copy the example-app directory (where the Dockerfile lives) into the container and make it the working directory.
# ADD ./app/app /

# # Now tell Docker what command to run when the container starts
# CMD ["/app"]

# # Set the PORT environment variable inside the container and expose it to the host so we can access our application
# EXPOSE $PORT

# # Examples to build and run
# # at project root
# # CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app/app app/app.go
# # CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -installsuffix netgo -o app/app app/app.go
# # CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -installsuffix netgo -installsuffix cgo -o app/app app/app.go
# # docker build ./ -f Dockerfile.scratch -t cep-provider-flat
# # docker run -it cep-provider-flat

# # https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker
# # https://medium.com/statuscode/golang-docker-for-development-and-production-ce3ad4e69673




# # Stage #1
# # The responsibility of this stage is to build an image which can build your Golang executable and extract the artefact.

# # Dockerfile.builder

# # # start a golang base image, version 1.8
# # FROM golang:1.8

# # #switch to our app directory
# # RUN mkdir -p /go/src/helloworld  
# # WORKDIR /go/src/helloworld

# # #copy the source files
# # COPY main.go /go/src/helloworld

# # #disable crosscompiling 
# # ENV CGO_ENABLED=0

# # #compile linux only
# # ENV GOOS=linux

# # #build the binary with debug information removed
# # RUN go build  -ldflags '-w -s' -a -installsuffix cgo -o helloworld  
# # To build it manually run this command to build it. 
# # $ docker build -f Dockerfile.builde -t builder:latest .

# # Copy the compiled artifact to your local disk 
# # $ docker container cp [id_of_container]:/go/src/helloworld/helloworld helloworld




# # Stage #2
# # The responsibility of this stage is to copy the artefact into the smallest possible image

# # Dockerfile.production

# # # start with a scratch (no layers)
# # FROM scratch

# # # copy our static linked library
# # COPY helloworld helloworld

# # # tell we are exposing our service on port 8080
# # EXPOSE 8080

# # # run it!
# # CMD ["./helloworld"]  
# # To build it manually run this command to build it. 
# # $ docker build -f Dockerfile.production -t helloworld:latest .