#!/bin/bash -x

PROJECT_DIR=../..

cd `dirname $0`

cp -v ${PROJECT_DIR}/dev.env.example ${PROJECT_DIR}/dev.env
mkdir -p ${PROJECT_DIR}/log/mysql
touch ${PROJECT_DIR}/log/mysql/mysqld.log
chmod 666 ${PROJECT_DIR}/log/mysql/mysqld.log
