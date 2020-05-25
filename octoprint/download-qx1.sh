#!/bin/sh
snapshot_number=$0
delay_seconds=$1
data_directory=$2
snapshot_directory=$3
snapshot_file_name=$4
snapshot_full_path=$5

url_file="$snapshot_file_name/$snapshot_full_path.url"

if [ -f $url_file ] ; then
	url=$(head -n 1 $url_file)

	mkdir -p $snapshot_directory
	mkdir -p $snapshot_file_name
	wget --quiet -O "$snapshot_file_name/$snapshot_full_path" $url

	rm $url_file
fi
