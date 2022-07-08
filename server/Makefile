GOOGLE_APPLICATION_CREDENTIALS ?= high-keel-345423-firebase-adminsdk-diqhx-fae01bf8e1.json

build.question:
	go build -o bin/question cmd/question/main.go
run.question.fb:
	GOOGLE_APPLICATION_CREDENTIALS=$(GOOGLE_APPLICATION_CREDENTIALS) \
	go run cmd/question/main.go
run.question:
	go run cmd/question/main.go	
test.local.qanda.create:
	grpcurl -plaintext -d '{"question":{"metadata":{"name":"string"},"author":{"id":"string"},"spec":{"content":{"text": {"value": "string"}}}}}' localhost:8080 questions.v1alpha1.QuestionService/CreateQuestion
test.local.qanda.list:
		grpcurl -plaintext localhost:8080 questions.v1alpha1.QuestionService/ListQuestions
test.local.qanda.get:
	grpcurl -plaintext -d '{ "id": "${ID}" }' localhost:8080 products.v1alpha1.ProductService/GetProduct
generate.swagger-ui:
	statik -f -src=../web/dist/swagger-ui/ -dest=internal/