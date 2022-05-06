FROM golang:1.15-alpine AS base

ENV GOPATH=
EXPOSE 5555
ADD . $project_dir

# install common dependencies
RUN go mod download
RUN apk add --no-cache curl python3 build-base


# build none-attack
FROM base AS none-attack
WORKDIR none-attack/
ENTRYPOINT ["go", "run", "main.go"]

# build JohnTheRipper binary and conf
FROM golang:1.15-alpine AS john-builder
WORKDIR /app
RUN apk add --no-cache git build-base perl
RUN git clone --depth 1 https://github.com/magnumripper/JohnTheRipper
RUN cd JohnTheRipper/src && ./configure --without-openssl && make -s clean && make -sj4

# build brute-force-secret
FROM base AS brute-force-secret
WORKDIR brute-force-secret/
COPY --from=john-builder /app/JohnTheRipper/run/ run/
ENTRYPOINT ["go", "run", "main.go"]
