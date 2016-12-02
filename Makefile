build:
	CGO_ENABLED=0 GOOS=linux go build -o pdf2text .

release: build
	docker build -t kszarek/pdf2text -f Dockerfile .

start:
	docker run -i --rm -p 5000:5000 --tmpfs /tmp kszarek/pdf2text
