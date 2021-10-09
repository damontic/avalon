all: publish

build:
	go build -o avalon
	tar czf avalon.tar.gz avalon html server

publish: build
	scp avalon.tar.gz damontic@${AVALON_SERVER_IP}:/home/damontic

clean:
	rm avalon.tar.gz avalon
