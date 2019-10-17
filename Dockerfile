FROM ubuntu:18.04-golang1.10.4 AS build
# Install tools required for project
# Run `docker build --no-cache .` to update dependencies
#RUN apk add --no-cache git
#RUN go get -u github.com/kardianos/govendor

# List project dependencies with Gopkg.toml and Gopkg.lock
# These layers are only re-built when Gopkg files are updated
#COPY Gopkg.lock Gopkg.toml /go/src/project/
WORKDIR /root/go/src/github.com/zhanglianx111/harbor-exporter

# Copy the entire project and build it
# This layer is rebuilt when a file changes in the project directory
COPY . /root/go/src/github.com/zhanglianx111/harbor-exporter
RUN go build -o /bin/harbor-exporter && chmod +x /bin/harbor-exporter


# This results in a single layer image
FROM ubuntu:18.04
COPY --from=build /bin/harbor-exporter /bin/harbor-exporter
COPY --from=build /root/go/src/github.com/zhanglianx111/harbor-exporter/config/config.yaml /etc/harbor-exporter/config.yaml
RUN apt-get update && apt-get install --no-install-recommends -y ca-certificates && rm -rf /var/lib/apt/lists/*

EXPOSE 9001

ENTRYPOINT ["/bin/harbor-exporter"]
