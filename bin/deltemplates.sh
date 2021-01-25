#!/bin/bash
# Delete all the templates in a given app
#
# !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
# WARNING: DO NOT RUN THIS AGAINST PRODUCTION APPLICATION
# !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
APP=$1

if [ "$1" == "" ]; then
	echo "Please enter an app name"
	exit 0
fi
./iotc.exe deviceTemplates list -a $APP --top 50 --format csv | awk -F, 'NR!=1 {if (NF>=3) {print $2}}' > dt.txt
echo ""
while read line
do
	echo "Deleting $line..."
	./iotc.exe deviceTemplates remove -a $APP --id $line
done < dt.txt

echo "Deleted `wc -l < dt.txt` templates"
rm dt.txt

