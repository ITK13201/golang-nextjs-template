#!/bin/bash -x

PROJECT_DIR=../..

cd `dirname $0`

# init environment variables
cp -v ${PROJECT_DIR}/dev.env.example ${PROJECT_DIR}/dev.env

# init mysql log files
mkdir -p ${PROJECT_DIR}/log/mysql
touch ${PROJECT_DIR}/log/mysql/mysqld.log

# init nginx log files
mkdir -p -p ${PROJECT_DIR}/log/nginx
touch ${PROJECT_DIR}/log/nginx/access.log
touch ${PROJECT_DIR}/log/nginx/error.log

# init log file permission
find ${PROJECT_DIR}/log -type f -print | xargs chmod 666
