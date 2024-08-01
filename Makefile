BINARY_PATH := ./bin/app
BINARY_PATH_WINDOWS := ./bin/app.exe
BINARY_MIGRATE_PATH := ./bin/migrate
MAIN_PATH := ./cmd/app/main.go
MIGRATE_PATH := ./cmd/migrate/main.go

build-windows:
	@echo 'Building golang'
	@CGO_ENABLED=1 GOARCH=amd64 GOOS=windows go build -o ${BINARY_PATH_WINDOWS} ${MAIN_PATH}	
	@echo 'Build completed'

build:
	@echo 'Building golang'
	@CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -o ${BINARY_PATH} ${MAIN_PATH}	
	@echo 'Build completed'

build-migrate:
	@echo 'Building migrate script'
	@CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -o ${BINARY_MIGRATE_PATH} ${MIGRATE_PATH}	
	@echo 'Build completed'

build-frontend:
	@echo 'Building frontend'
	@cd frontend && npm run build
	@echo 'Build completed'

run-prod: build-frontend build
	@echo 'Starting server'
	@${BINARY_PATH}

run: build 
	@echo 'Starting server'
	@${BINARY_PATH}

migrate: build-migrate
	@echo 'Launch migrate script'
	@${BINARY_MIGRATE_PATH} --run=create
	@echo 'End migrating'

migrate-reset: build-migrate
	@echo 'Launch migrate script'
	@${BINARY_MIGRATE_PATH} --run=reset
	@echo 'End migrating'

migrate-drop: build-migrate
	@echo 'Launch migrate script'
	@${BINARY_MIGRATE_PATH} --run=drop
	@echo 'End migrating'

