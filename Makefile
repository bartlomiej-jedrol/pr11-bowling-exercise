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

build_cli_linux:
	GOOS=linux GOARCH=amd64 go build -o build/bowling-linux cmd/cli/main.go

build_cli_windows:
	GOOS=windows GOARCH=amd64 go build -o build/bowling-windows.exe cmd/cli/main.go

build_cli_mac:
	GOOS=darwin GOARCH=amd64 go build -o build/bowling-mac cmd/cli/main.go

build_clis: build_cli_linux build_cli_windows build_cli_mac
	chmod +x build/bowling-linux build/bowling-mac
