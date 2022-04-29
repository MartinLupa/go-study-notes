# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

# To make things easier when running the rest of our commands, let’s create a directory inside the image that we are building. 
# This also instructs Docker to use this directory as the default destination for all subsequent commands. 
# This way we do not have to type out full file paths but can use relative paths based on this directory.
WORKDIR /app

# Usually the very first thing you do once you’ve downloaded a project written in Go is to install the modules necessary to compile it.
# But before we can run go mod download inside our image, we need to get our go.mod and go.sum files copied into it.
COPY go.mod ./
COPY go.sum ./

# Now that we have the module files inside the Docker image that we are building, we can use the RUN command to execute the command go mod download there as well.
RUN go mod download

# Copy our source code into the image.
# This COPY command uses a wildcard to copy all files with .go extension located in the current directory on the host (the directory where the Dockerfile is located) 
# into the current directory inside the image.
COPY *.go ./

# The result of this command will be a static application binary named docker-gs-ping and located in the root of the filesystem of the image that we are building.
RUN go build -o /docker-gs-ping

# Now, all that is left to do is to tell Docker what command to execute when our image is used to start a container.
CMD ["/docker-gs-ping"]

# TO CREATE THE IMAGE, run: docker build --tag docker-ps-ping .
# The docker build command creates Docker images from the Dockerfile and a “context”. A build context is the set of files located in the specified path or URL. 
# The Docker build process can access any of the files located in the context.
# The build command optionally takes a --tag flag. This flag is used to label the image with a string value, which is easy for humans to read and recognise. 
# If you do not pass a --tag, Docker will use latest as the default value.

# IMAGE TAG:
# docker image tag docker-gs-ping:latest docker-gs-ping:v1.0
# Changes the tag from latest to v1.0.

# RUN CONTAINER:
# To publish a port for our container, we’ll use the --publish flag (-p for short) on the docker run command. The format of the --publish command is [host_port]:[container_port]. 
# So if we wanted to expose port 8080 inside the container to port 3000 outside the container, we would pass 3000:8080 to the --publish flag.
# docker run --publish 8080:8080 docker-gs-ping