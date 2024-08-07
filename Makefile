dockerize:
	docker build -t corentus .
	docker run -d -p 8080:8080 --name corentus-container corentus
