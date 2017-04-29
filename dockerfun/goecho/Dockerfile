FROM scratch
ENV NAME goecho
ENV APP /app
ADD ./builds/${NAME} ${APP}
EXPOSE 25000
ENTRYPOINT [ "/app" ]
