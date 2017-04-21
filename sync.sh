#!/bin/bash
set -e

FILES=`ls ~/.ssh/ -I authorized_keys -I known_hosts`
for file in $FILES ; do
	name=`basename $file`
	dest="Secure Notes\\SSH/$name"
	echo "Syncing $file to $dest"
	cat ~/.ssh/$file  | lpass edit --non-interactive  --notes "$dest"
done

echo "Syncing..."
lpass sync

UPSTREAM=`lpass ls "Secure Notes\\SSH" --color=never | tr -d ' '`
for entry in $UPSTREAM; do
	file=`echo $entry | cut -d '/' -f 2 | cut -d '[' -f 1`
	 if [ ! -f ~/.ssh/$file ] && [ ! -d ~/.ssh/$file ]; then
	    echo "Creating ~/.ssh/$file"
	    id=`echo $entry | cut -d ":" -f 2 | tr -d '[]'`
	    lpass show "$id" --notes >  ~/.ssh/$file
	fi	
done	

# lpass ls "Secure Notes\\SSH"

# echo "Failures....."
# lpass ls "Secure Notes\\SSH" --color=never | grep -w "\b[id: 0]\b"

