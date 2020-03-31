# Initialize a new build stage, rename it and set Golang as base image for subsequent instructions
FROM golang:latest AS buildStageCompiling

# Copy source files from host's context into the container's (new created) workspace
WORKDIR /go/src/pv
COPY . .

# Disable cgo to avoid error "standard_init_linux.go:211: exec user process caused "no such file or directory""
ENV CGO_ENABLED=0

# Downloat all imported packages and compile the program
RUN go get -d -v ./...
RUN go build -a

# Copy the executable into an empty image and execute it with container start
FROM scratch AS buildStageRuntime
COPY --from=buildStageCompiling /go/src/pv/pv /
EXPOSE 80/tcp
ENTRYPOINT ["/pv"]