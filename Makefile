up:
	cp .air.normal.conf .air.conf
	docker-compose up --build

debug:
	cp .air.debug.conf .air.conf
	docker-compose up --build

down:
	docker-compose down

stop:
	docker-compose stop

del:
	docker system prune -a

db:
	docker-compose exec db mysql  -u root -p go_mission

