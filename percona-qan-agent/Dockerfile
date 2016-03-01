FROM ubuntu:latest

WORKDIR /opt

RUN apt-get -y update && apt-get install -y apt-transport-https mysql-client

ADD https://www.percona.com/redir/downloads/TESTING/ppl/open-source/ppl-agent.tar.gz /opt 
ADD run.sh /opt

RUN tar zxf ppl-agent.tar.gz --strip-components=1

CMD ["/opt/run.sh"]
