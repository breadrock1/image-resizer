FROM golang:1.22 as build

ENV BIN_FILE /opt/resizer/resizer
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS
RUN CGO_ENABLED=0 go build \
        -ldflags "$LDFLAGS" \
        -o ${BIN_FILE} cmd/resizer/*

FROM alpine:3.9

ENV BIN_FILE /opt/resizer/resizer
COPY --from=build ${BIN_FILE} ${BIN_FILE}

ENV CONFIG_FILE /etc/resizer/config.toml
COPY ./configs/config.toml ${CONFIG_FILE}

CMD ${BIN_FILE} -c ${CONFIG_FILE}
