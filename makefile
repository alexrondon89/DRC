.PHONY: facebook-collector pg-services csv-generator all

pg-services:
	chmod +x ./scripts/services.sh
	./scripts/services.sh

facebook-collector: pg-services
	chmod +x ./scripts/collector.sh
	./scripts/collector.sh

csv-generator: pg-services
	@read -p "Indica la ruta absoluta donde quieres guardar el documento .csv: " OUTPUT_PATH; \
	read -p "Seleccione el tipo de informacion a recuperar: " TYPE; \
	chmod +x ./scripts/csv-generator.sh; \
	./scripts/csv-generator.sh "$$OUTPUT_PATH" "$$TYPE"

all: pg-services facebook-collector csv-generator
