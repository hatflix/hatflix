db: ## Subir o banco localmente.
	docker-compose -f ./build/db/docker-compose.yml up -d

fresh-db: ## Sobe o banco localmente mas recriando tudo do zero
	docker-compose -f ./build/db/docker-compose.yml down -v
	docker-compose -f ./build/db/docker-compose.yml up --build --remove-orphans -d
