#
# Copyright (c) 2012-2015 Codenvy, S.A.
# All rights reserved. This program and the accompanying materials
# are made available under the terms of the Eclipse Public License v1.0
# which accompanies this distribution, and is available at
# http://www.eclipse.org/legal/epl-v10.html
#
# Contributors:
#   Codenvy, S.A. - initial API and implementation
#

FROM codenvy/go

EXPOSE 8080
ENV CODENVY_APP_PORT_8080_HTTP 8080

ENV GOPATH /home/user/application

RUN mkdir -p /home/user/application/src

ENV CODENVY_APP_BIND_DIR /home/user/application/src

VOLUME ["/home/user/application/src"]

ENV DEBIAN_FRONTEND noninteractive
  
# Install MYSQL
RUN sudo -E bash -c "apt-get -y --no-install-recommends install mysql-server" && \
    sudo sed -i.bak 's/127.0.0.1/0.0.0.0/g' /etc/mysql/my.cnf && \
    sudo service mysql restart && \
    sudo mysql -uroot -e "GRANT ALL PRIVILEGES ON *.* TO 'root'@'%'; FLUSH PRIVILEGES;"
 
EXPOSE 3306
 
# Create MySQL user and database, and start it
RUN sudo service mysql restart && \
    CODENVY_MYSQL_PASSWORD=password && \
    CODENVY_MYSQL_DB=test && \
    CODENVY_MYSQL_USER=test && \
    echo "MySQL password: $CODENVY_MYSQL_PASSWORD" >> /home/user/.mysqlrc && \
    echo "MySQL user    : $CODENVY_MYSQL_USER" >> /home/user/.mysqlrc && \
    echo "MySQL Database:$CODENVY_MYSQL_DB" >> /home/user/.mysqlrc && \
    sudo mysql -uroot -e "CREATE USER '$CODENVY_MYSQL_USER'@'%' IDENTIFIED BY '"$CODENVY_MYSQL_PASSWORD"'" && \
    sudo mysql -uroot -e "GRANT ALL PRIVILEGES ON *.* TO '$CODENVY_MYSQL_USER'@'%' IDENTIFIED BY '"$CODENVY_MYSQL_PASSWORD"'; FLUSH PRIVILEGES;" && \
    sudo mysql -uroot -e "CREATE DATABASE $CODENVY_MYSQL_DB;"
 
CMD sudo service mysql start > /dev/null

WORKDIR /home/user/application/src

# 1. Get dependencies from source code of application
# 2. Start application
CMD go get -d && go run $executable:-main.go$