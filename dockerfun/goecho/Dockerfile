FROM scratch
MAINTAINER daved <disjuncter@gmail.com>
ENV NAME goecho
ENV SRV /srv
ADD ./builds/${NAME} ${SRV}/
EXPOSE 25000
ENTRYPOINT ["/srv/goecho"]
