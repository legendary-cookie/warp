all:
	go build

run: all
	./warp

service: all
	systemctl restart warp.service
