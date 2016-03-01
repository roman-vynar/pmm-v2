FROM ubuntu:latest

WORKDIR /opt

RUN apt-get -y update && apt-get install -y apt-transport-https mysql-server python supervisor

ADD https://www.percona.com/downloads/TESTING/ppl/open-source/ppl-datastore.tar.gz https://www.percona.com/downloads/TESTING/ppl/open-source/ppl-qan-app.tar.gz /opt/
ADD supervisord.conf /etc/supervisor/conf.d/qan.conf
ADD install-qan.sh /opt

RUN mkdir qan-api && tar zxf ppl-datastore.tar.gz --strip-components=1 -C qan-api
RUN mkdir qan-app && tar zxf ppl-qan-app.tar.gz   --strip-components=1 -C qan-app
RUN /opt/install-qan.sh

EXPOSE 9001 9002

CMD ["supervisord"]
