## Percona Monitoring and Management project

### Requirements

    easy_install pip
    pip install docker-compose

SELinux should be off or permissive due to Grafana issue.

### Monitor setup

    cd monitor
    docker-compose up -d

Add host for monitoring at any time:

    ./pmm-admin -add db1 192.168.56.107

### DB node setup

    cd dbnode
    <Edit my.cnf with MySQL credentials>
    docker-compose up -d

MySQL privileges:

    GRANT SELECT, UPDATE, DELETE, DROP ON performance_schema.* TO 'pmm'@'localhost';
    GRANT REPLICATION CLIENT, PROCESS, SUPER, SELECT ON *.* TO 'pmm'@'localhost' IDENTIFIED BY 'percona2016';

### Percona Query Analytics

Application for monitor host:

    cd percona-qan
    <Edit docker-compose.yml and set HOST_IP to host IP>
    docker-compose up -d

Agent for db node:

    cd percona-qan-agent 
    <Edit my.cnf with root(!) MySQL credentials>
    <Edit docker-compose.yml and set HOST_IP to monitor IP and update path to slow.log>
    docker-compose up -d

### Credits

Roman Vynar, Daniel Nichter, Kenny Gryp, Ben Mildren.
