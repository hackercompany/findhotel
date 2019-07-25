default:fmt clean
	$(MAKE) build

build: main_application

main_application:
	go build -o bin/main

fmt:
	go fmt ./src/...

clean:
	rm -rf `find ./vendor/src -type d -name .git` \
	&& rm -rf `find ./vendor/src -type d -name .hg` \
	&& rm -rf `find ./vendor/src -type d -name .bzr` \
	&& rm -rf `find ./vendor/src -type d -name .svn`
	rm -rf ./bin/*