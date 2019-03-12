# HZDeit

A progressive platform for rising an awerness about food consumption cost.
Provided as Platform-as-a-Service.

## Usage
Frontend and backend are docerised.
Folders contain Dockerfile.

$: docker build -t TAG:V .
where TAG - is name for your container
    V - version

Solution is cloud-native.
It contains kubernetes config files.
You can easely deploy it in your cloud.

$: kubectl apply -f "/path/to/config.yml"

## Technologies
Backend
    - Go
    - Docker
    - Mysql
    - Kubernetes

Frontend:
    - HTML
    - CSS
    - JS
    - Framework7
