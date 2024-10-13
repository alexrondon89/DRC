.PHONY: facebook-collector pg-services csv-generator all

pg-services:
	docker compose -f ./information-collector-service/docker-compose.yaml up -d

facebook-collector: pg-services
	chmod +x ./scripts/collector.sh
	./scripts/collector.sh

csv-generator: pg-services
	@read -p "indica la ruta absoluta donde quieres guardar el documento .csv: " PATH;\
	chmod +x ./scripts/csv-generator.sh
	./scripts/csv-generator.sh $$PATH

all: pg-services facebook-collector csv-generator
