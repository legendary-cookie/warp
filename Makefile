all:
	go build
	
service: all
	systemctl restart warp.service
