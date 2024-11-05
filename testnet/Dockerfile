FROM golang:1.22 AS op-node

WORKDIR /app

ENV REPO=https://github.com/mdehoog/optimism.git
ENV VERSION=l3
ENV COMMIT=445adf498c862eaef12e650909938d2abad9d143
RUN git clone $REPO --branch $VERSION --single-branch . && \
    git switch -c branch-$VERSION && \
    bash -c '[ "$(git rev-parse HEAD)" = "$COMMIT" ]'

RUN cd op-node && \
    make VERSION=$VERSION op-node


FROM golang:1.22 AS op-enclave

WORKDIR /app

COPY op-enclave/go.mod op-enclave/go.sum op-enclave/
RUN cd op-enclave && go mod download

COPY go.mod go.sum ./
RUN go mod download

COPY bindings/ bindings/
COPY op-enclave/ op-enclave/
COPY op-batcher/ op-batcher/
COPY op-da/ op-da/
COPY op-proposer/ op-proposer/

RUN cd op-enclave && go build -o ../bin/op-enclave ./cmd/enclave
RUN go build -o bin/op-batcher ./op-batcher/cmd
RUN go build -o bin/op-proposer ./op-proposer/cmd
RUN go build -o bin/op-da ./op-da/cmd


FROM golang:1.22 AS op-geth

WORKDIR /app

ENV REPO=https://github.com/mdehoog/op-geth.git
ENV VERSION=witness2
ENV COMMIT=d8f4cde5a85265a258dfcd21f0c735d8aefaa07d
RUN git clone $REPO --branch $VERSION --single-branch . && \
    git switch -c branch-$VERSION && \
    bash -c '[ "$(git rev-parse HEAD)" = "$COMMIT" ]'

RUN go run build/ci.go install -static ./cmd/geth


FROM ubuntu:22.04 AS op-reth

WORKDIR /app

RUN apt-get update && apt-get -y upgrade && apt-get install -y git libclang-dev pkg-config curl build-essential
RUN curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh -s -- -y

ENV PATH="/root/.cargo/bin:${PATH}"

ENV REPO=https://github.com/paradigmxyz/reth.git
ENV VERSION=v1.1.0
ENV COMMIT=1ba631ba9581973e7c6cadeea92cfe1802aceb4a
RUN git clone $REPO --branch $VERSION --single-branch . && \
    git switch -c branch-$VERSION && \
    bash -c '[ "$(git rev-parse HEAD)" = "$COMMIT" ]'

RUN cargo build --bin op-reth --features jemalloc,asm-keccak,optimism --manifest-path crates/optimism/bin/Cargo.toml


FROM ubuntu:22.04

RUN apt-get update && apt-get install -y curl jq

WORKDIR /app

COPY --from=op-node /app/op-node/bin/op-node ./
COPY --from=op-geth /app/build/bin/geth ./
COPY --from=op-reth /app/target/debug/op-reth ./reth
COPY --from=op-enclave /app/bin/op-enclave ./
COPY --from=op-enclave /app/bin/op-batcher ./
COPY --from=op-enclave /app/bin/op-proposer ./
COPY --from=op-enclave /app/bin/op-da ./
COPY testnet/entrypoint-batcher.sh ./
COPY testnet/entrypoint-da.sh ./
COPY testnet/entrypoint-enclave.sh ./
COPY testnet/entrypoint-geth.sh ./
COPY testnet/entrypoint-node.sh ./
COPY testnet/entrypoint-proposer.sh ./
COPY testnet/entrypoint-reth.sh ./
COPY deployments/ deployments/
