# Development
FROM mcr.microsoft.com/devcontainers/go:0-1.19 as dev


# ---------------------------------------------------------------------


# Deploy Builder
FROM mcr.microsoft.com/devcontainers/go:0-1.19 as deploy-builder

ENV ROOT=/go/src/app

WORKDIR ${ROOT}/src

COPY . ${ROOT}
RUN go build -trimpath -ldflags "-w -s" -o main


# --------------------------------------------------------------------


# Deploy
FROM debian:bullseye-slim as deploy

ENV ROOT=/go/src/app

RUN apt-get update

COPY --from=deploy-builder ${ROOT}/src/main .

CMD [ "main" ]
