GOOGLE_APPLICATION_CREDENTIALS ?= high-keel-345423-firebase-adminsdk-diqhx-fae01bf8e1.json


build.question:
	go build -o bin/question cmd/question/main.go
run.question:
	GOOGLE_APPLICATION_CREDENTIALS=$(GOOGLE_APPLICATION_CREDENTIALS) go run cmd/question/main.go

generate.swagger-ui:
	statik -f -src=../web/dist/swagger-ui/ -dest=internal/