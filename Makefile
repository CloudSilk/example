run:
	SERVICE_MODE=ALL  go run main.go --ui ./web/dist --port 48080
gen-doc:
	swag init --parseDependency --parseInternal --parseDepth 1

start-web:
	WEB_BASE=/web yarn start
build-web:
	cd web && WEB_BASE=/web yarn build