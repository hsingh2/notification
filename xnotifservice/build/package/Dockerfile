ARG BASE_IMAGE=dockerhub.cisco.com/vms-platform-dev-docker/vms-base-stretch:3.8.0-370

###
### Build container
###

FROM golang:1.12-stretch AS builder
ARG BUILD_FLAGS
ENV BUILDER_FLAGS -mod=vendor
ADD . /app/
WORKDIR /app
RUN make dist

###
### Distribution Container
###

FROM ${BASE_IMAGE}
EXPOSE 9213
WORKDIR /var/run/xnotif/
CMD ["/usr/bin/xnotif"]
COPY --from=builder /app/dist/root/ /
