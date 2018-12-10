test:
	go test -v ./...

run:
	go run cmd/gophrase.go

build:
	go build cmd/gophrase.go

binaries_files:
	set -eux \
		&& export BUILD_NUMBER='v0.0.1' \
		&& export DEPLOY_FILENAME='cmd/gophrase.go' \
		&& export applicationName='gophrase'  \
		&& echo "Building the Linux x64 binary" \
		&& GOOS=linux GOARCH=amd64 go build -v -o binaries/amd64/linux/$${applicationName}-$${BUILD_NUMBER}.linux.amd64 $${DEPLOY_FILENAME} \
		&& echo "Building the Mac x64 binary" \
		&& GOOS=darwin GOARCH=amd64 go build -v -o binaries/amd64/darwin/$${applicationName}-$${BUILD_NUMBER}.darwin.amd64 $${DEPLOY_FILENAME} \
		&& echo "Building the Windows x64 binary" \
		&& GOOS=windows GOARCH=amd64 go build -v -o binaries/amd64/windows/$${applicationName}-$${BUILD_NUMBER}.windows.amd64.exe $${DEPLOY_FILENAME} \
		&& echo "Done."
#		&& echo "Packaging static files" \
#		&& /go/bin/packr --verbose --input $${DEPLOY_FILENAME} \
#		&& /go/bin/packr clean
