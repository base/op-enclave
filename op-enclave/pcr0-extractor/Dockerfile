# Build skopeo
FROM golang@sha256:367bb5295d3103981a86a572651d8297d6973f2ec8b62f716b007860e22cbc25 AS skopeo
WORKDIR /app
ENV REPO=https://github.com/containers/skopeo.git
# v1.12.0
ENV COMMIT=a55290973794d93f602a027e795bf510bd3cad01
RUN git init && \
    git remote add origin $REPO && \
    git fetch --depth=1 origin $COMMIT && \
    git reset --hard FETCH_HEAD
RUN CGO_ENABLED=0 DISABLE_DOCS=1 make BUILDTAGS=containers_image_openpgp GO_DYN_FLAGS=

# Build umoci
FROM golang@sha256:367bb5295d3103981a86a572651d8297d6973f2ec8b62f716b007860e22cbc25 AS umoci
WORKDIR /app
ENV REPO=https://github.com/opencontainers/umoci.git
# v0.4.7
ENV COMMIT=17f38511d61846e2fb8ec01a1532f3ef5525e71d
RUN git init && \
    git remote add origin $REPO && \
    git fetch --depth=1 origin $COMMIT && \
    git reset --hard FETCH_HEAD
RUN go build -o bin/umoci ./cmd/umoci

# Get the EIF
FROM golang@sha256:367bb5295d3103981a86a572651d8297d6973f2ec8b62f716b007860e22cbc25 AS op-enclave
WORKDIR /app
COPY --from=skopeo /app/bin/skopeo /app/bin/skopeo
COPY --from=umoci /app/bin/umoci /app/bin/umoci

ENV TAG=v0.0.1-rc5
RUN bin/skopeo --insecure-policy copy docker://ghcr.io/base/op-enclave:$TAG oci:op-enclave:latest
RUN bin/umoci unpack --image op-enclave bundle

# Extract PCR0
FROM amazonlinux:2
RUN amazon-linux-extras enable aws-nitro-enclaves-cli && \
    yum clean metadata && \
    yum update -y && \
    yum install -y aws-nitro-enclaves-cli aws-nitro-enclaves-cli-devel jq

COPY --from=op-enclave /app/bundle/rootfs/build/eif.bin /app/eif.bin
COPY extract-pcr0.sh /extract-pcr0.sh
RUN chmod +x /extract-pcr0.sh

ENTRYPOINT ["/extract-pcr0.sh"] 