#! /bin/sh
sudo docker stop mysql-db
sudo docker run --name mysql-db -d -p 3306:3306 mysql-db