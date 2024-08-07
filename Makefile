dockerize:
	docker build -t corentus .
	docker run -d -p 4000:4000 --name corentus-container corentus
