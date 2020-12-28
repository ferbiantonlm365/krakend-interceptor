.PHONY: plugins

all: plugins

plugins:
	go build -buildmode=plugin -o plugins/router-plugin.so ./router-plugin
	go build -buildmode=plugin -o plugins/proxy-plugin.so ./proxy-plugin