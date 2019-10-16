FROM golang:1.12-alpine AS build

# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
#RUN apk add --no-cache git
#RUN go get -u github.com/kardianos/govendor

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
#COPY Gopkg.lock Gopkg.toml /go/src/project/
WORKDIR /go/src/project/
# Install library dependencies
#RUN dep ensure -vendor-only

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY . /go/src/project/
RUN go build -o /bin/harbor-exporter

# This results in a single layer image
FROM scratch
COPY --from=build /bin/harbor-exporter /bin/harbor-exporter
COPY  --from=build /go/src/project/harbor-exporter/config/config.yaml /etc/harbor-exporter/config.yaml
EXPOSE 9001

ENTRYPOINT ["/bin/harbor-exporter"]
