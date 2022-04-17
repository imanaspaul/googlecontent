dev:
	air
db:
	docker-compose -f docker-composepg.yml up -d
build:
	go build -o