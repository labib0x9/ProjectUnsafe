.PHONY: backend frontend run

backend:
	go run main.go

frontend:
	cd ../PUF-CLAUDE-GIT && npm run dev

run:
	@echo "Starting postgres and redis"
	brew services start postgresql & \
	brew services start redis & \
	sleep 2 && \
	@echo "Starting backend and frontend..." \
	@trap 'kill 0' SIGINT; \
	make frontend & \
	make backend & \
	wait