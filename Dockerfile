# dynamic config
ARG             BUILD_DATE
ARG             VCS_REF
ARG             VERSION

# build
FROM            golang:1.17.5-alpine as builder
RUN             apk add --no-cache git gcc musl-dev make
ENV             GO111MODULE=on
WORKDIR         /go/src/moul.io/wipchat
COPY            go.* ./
RUN             go mod download
COPY            . ./
RUN             make install

# minimalist runtime
FROM alpine:3.14.0
LABEL           org.label-schema.build-date=$BUILD_DATE \
                org.label-schema.name="wipchat" \
                org.label-schema.description="" \
                org.label-schema.url="https://moul.io/wipchat/" \
                org.label-schema.vcs-ref=$VCS_REF \
                org.label-schema.vcs-url="https://github.com/moul/wipchat" \
                org.label-schema.vendor="Manfred Touron" \
                org.label-schema.version=$VERSION \
                org.label-schema.schema-version="1.0" \
                org.label-schema.cmd="docker run -i -t --rm moul/wipchat" \
                org.label-schema.help="docker exec -it $CONTAINER wipchat --help"
COPY            --from=builder /go/bin/wipchat /bin/
ENTRYPOINT      ["/bin/wipchat"]
#CMD             []
