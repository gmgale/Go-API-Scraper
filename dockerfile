FROM golang:latest

# Copy the local package files to the container’s workspace.
ADD . \Go\src\github.com\gmgale\result_task
# Install our dependencies
RUN go get http://github.com\gorilla\mux

# Install api binary globally within container 
RUN go install http://github.com\gmgale\result_task

# Set binary as entrypoint
ENTRYPOINT \Go\bin\

# Expose default port (8080)
EXPOSE 8080