db: ## Subir o banco localmente.
	docker-compose -f ./build/db/docker-compose.yml up -d

fresh-db: ## Sobe o banco localmente mas recriando tudo do zero
	docker-compose -f ./build/db/docker-compose.yml down -v
	docker-compose -f ./build/db/docker-compose.yml up --build --remove-orphans -d

gqlgen: ## Gera os arquivos para a API GraphQL
	go generate ./config/http/config.go

run:
	go generate ./cmd/http/http.go