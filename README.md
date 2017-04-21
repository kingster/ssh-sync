# SSH key(s) management with LastPass 
 
This repository has a tool called `sshsync` which uses LastPass (secure notes)
to backup and restore ssh key files. 

## How to install

Fist install the LastPass command line client (`lpass`) from a `lastpass-cli`
package:

	brew install lastpass-cli --with-pinentry

Next install this package. 
	
	git clone https://github.com/kingster/ssh-sync.git
	make install

## How to use

Run:

	sshsync

It will sync the ~/.ssh folder with  "Secure Notes/SSH" folder.

