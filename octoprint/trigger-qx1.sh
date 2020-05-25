#!/bin/sh
snapshot_number=$0
delay_seconds=$1
data_directory=$2
snapshot_directory=$3
snapshot_file_name=$4
snapshot_full_path=$5

mkdir -p $snapshot_directory
mkdir -p $snapshot_file_name

/home/pi/devel/git/golang-sony-remote-api-test/sony-api \
  --action actTakePicture \
  --endpoint http://192.168.122.1:8080 \
  --output_picture_url "$snapshot_file_name/$snapshot_full_path.url"

# Uncomment this to download the image in the same step.
# This will enable you to see the preview in the Octolapse UI.
#
# However, this is downloading the image from the network and takes some time.
# Carriage might end up being stopped for too long, affecting the quality
# of your 3D Print.
# /home/pi/devel/git/golang-sony-remote-api-test/octoprint/download-qx1.sh $1 $2 $3 $4 $5
