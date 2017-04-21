#!/bin/bash
set -e

for file in `ls ~/.ssh/ -I authorized_keys -I known_hosts -I *.bak`
do
  name=`basename $file`
  dest="Secure Notes\\SSH/$name"
  echo "Syncing $file to $dest"
  cat ~/.ssh/$file  | lpass edit --non-interactive  --notes "$dest"
done

echo "Syncing..."
lpass sync

lpass ls "Secure Notes\\SSH"

echo "Failures....."

lpass ls "Secure Notes\\SSH" --color=never | grep -w "\b[id: 0]\b"

