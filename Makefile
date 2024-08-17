
prepare-npm:
	brew install npm

prepare-ganache-cli: prepare-npm
	npm install -g ganache-cli


run-host:
	docker-compose up -d

stop-host:
	docker-compose down

install:
	go install .