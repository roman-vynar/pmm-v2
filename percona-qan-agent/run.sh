#!/usr/bin/env bash

set -eu

cd /opt
if [ ! -e .my.cnf ]; then
   echo "You need to share .my.cnf to the container's /opt/.my.cnf"
   exit
fi
MYSQL_USER=$(cat .my.cnf|egrep "^user"|awk -F'=' '{print $2}')
MYSQL_PASSWORD=$(cat .my.cnf|egrep "^password"|awk -F'=' '{print $2}')

./install -user=$MYSQL_USER -password=$MYSQL_PASSWORD -host=127.0.0.1 $HOST_IP
tail -f /usr/local/percona/agent/percona-agent.log

