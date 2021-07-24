all:
	go build

run: all
	./warp

service: all
	systemctl restart warp.service


install: all
	mkdir -pv /etc/warp
	mkdir -pv /var/lib/warp
	cp template-configs/config.yml.tmpl /etc/warp/config.yml
	cp template-configs/mime.yml.tmpl /etc/warp/mime.yml
	cp template-configs/warp.service.tmpl /etc/systemd/system/warp.service
	cp warp /usr/bin/warp

uninstall:
	rm /usr/bin/warp
	systemctl disable warp.service
	systemctl stop warp.service
	rm /etc/systemd/system/warp.service
