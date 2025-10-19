FROM golang:1.24.6-bullseye

ARG TARGETARCH
ENV DEBIAN_FRONTEND=noninteractive

# Install curl, unzip
#RUN apt-get update && apt-get install -y --no-install-recommends \
#    curl unzip


RUN groupadd --gid 1000 example && \
    useradd --uid 1000 --gid example --create-home example

# Set ownership of directories needed later
RUN mkdir -p /opt/nsa2/example \
    && chown -R example:example /opt/nsa2/example

USER example

WORKDIR /opt/nsa2/example

ENV SF_WORKDIR=/opt/nsa2/example



COPY --chown=example:example cmd/ cmd/
COPY --chown=example:example internal/ internal/


COPY --chown=example:example go.mod go.sum ./
RUN go mod download

#COPY --chown=example:example . .

RUN go build -o otel-go-example cmd/main.go

EXPOSE 8080

CMD ["./otel-go-example"]
#ENTRYPOINT ["./entrypoint.sh"]