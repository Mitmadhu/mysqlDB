FROM mysql

ENV MYSQL_DATABASE=microservice
ENV MYSQL_USER=madhu
ENV MYSQL_PASSWORD=mad.ayush
ENV MYSQL_ROOT_PASSWORD=mad.ayush

# RUN --default-authentication-plugin=mysql_native_password

EXPOSE 3306