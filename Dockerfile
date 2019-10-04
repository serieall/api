FROM debian

WORKDIR /usr/local/bin

ADD ./bin/serieall-api /usr/local/bin/serieall-api
ADD VERSION /usr/local/bin/VERSION

ENTRYPOINT ["serieall-api"]
