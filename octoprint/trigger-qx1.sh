#!/bin/sh
snapshot_number=$0
delay_seconds=$1
data_directory=$2
snapshot_directory=$3
snapshot_file_name=$4
snapshot_full_path=$5

mkdir -p $3
mkdir -p $4

/home/pi/devel/git/golang-sony-remote-api-test/sony-api \
  --endpoint http://192.168.122.1:8080 \
  --action actTakePicture \
  --output_picture "$snapshot_file_name/$snapshot_full_path"
