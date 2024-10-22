# Usa la imagen base de Python 3.9
FROM python:3.9-slim

# Establecer el directorio de trabajo
WORKDIR /app

# Copia los archivos de requisitos en el contenedor
COPY requirements.txt .

# Instala las dependencias del sistema necesarias
RUN apt-get update && apt-get install -y \
    gcc \
    libatlas-base-dev \
    && rm -rf /var/lib/apt/lists/*

# Instala las dependencias de Python
RUN pip install --no-cache-dir -r requirements.txt

# Copia el resto de la aplicación al contenedor
COPY . .

# Comando por defecto para ejecutar la aplicación
CMD ["python", "cmd/main.py"]
