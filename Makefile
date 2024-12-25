.PHONY: run stop down


run:
	docker compose up -d

stop:
	docker stop postgres-labs

down:
	docker compose down
