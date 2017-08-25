# SSH key(s) management with LastPass 
 
This repository provides the tool called `ssh-sync-keys` which uses LastPass (secure notes)
to backup and restore ssh key files. 

## How to install

- First install the LastPass command line client (`lpass`) from a `lastpass-cli` package:

	```
	brew install lastpass-cli --with-pinentry
	```
- [Sign up on lastpass](https://lastpass.com/f?207276) if you dont have an account. This account will securely store the keys
- Next install this package. 

	```
	git clone https://github.com/kingster/ssh-sync.git
	make
	make install
	```

## How to use

Run:

	lpass login <your-email-address>
	ssh-sync-keys
	
	 

It will sync the ~/.ssh folder with  "Secure Notes/SSH" folder.

