sshsync: sync.sh
   
install:
	cp sync.sh  /usr/local/bin/sshsync

purge: 
	rm /usr/local/bin/sshsync