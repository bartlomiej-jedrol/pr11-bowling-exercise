AWS_LAMBDA_FUNCTION = "pr11-lambda"
AWS_LAMBDA_SOURCE_PATH = "cmd/lambda/main.go"
AWS_LAMBDA_BUILD_ZIP_PATH = "build/main.zip"

clean:
	rm -f bootstrap $(AWS_LAMBDA_BUILD_ZIP_PATH)

build: clean
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap $(AWS_LAMBDA_SOURCE_PATH)
	zip $(AWS_LAMBDA_BUILD_ZIP_PATH) bootstrap

push: build
	aws lambda update-function-code --function-name $(AWS_LAMBDA_FUNCTION) \
	--zip-file fileb://$(AWS_LAMBDA_BUILD_ZIP_PATH)