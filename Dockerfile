FROM golang:1.8.1
ENV APP /app
WORKDIR /go/src${APP}/
COPY *.go .
RUN CGO_ENABLED=0 GOOS=linux; go build -a -tags netgo -ldflags '-w' -o ${APP}

FROM scratch
ENV CERTS /etc/ssl/certs/ca-certificates.crt
COPY --from=0 ${CERTS} ${CERTS}
COPY --from=0 ${APP} ${APP}
EXPOSE 25000
ENTRYPOINT [ "/app" ]
