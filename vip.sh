#!/bin/bash

set -x
# Stripts off whitespace, no idea how to really read stdin it seems
VIPFILE=$(mktemp vipXXXXX)
while read input; do
	echo "$input" >> $VIPFILE;
done

vim -i "NONE" \
	-u "NONE" \
	$@ \
	-c "wq!" -- \
	$VIPFILE

cat $VIPFILE
rm -f $VIPFILE
