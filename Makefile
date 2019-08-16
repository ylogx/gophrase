test:
	go test -v ./...

run:
	go run cmd/gophrase.go

build: #setup_repo
	go build cmd/gophrase.go

install: build
	install -pv gophrase ~/.local/bin

setup_repo:
	mkdir -p ~/.cache/go/src/github.com/ylogx/
	ln -sf $$(pwd) ~/.cache/go/src/github.com/ylogx/

build_binaries:
	set -eux \
		&& export BUILD_NUMBER='v0.0.1' \
		&& export DEPLOY_FILENAME='cmd/gophrase.go' \
		&& export applicationName='gophrase'  \
		&& echo "Building the Linux x64 binary" \
		&& GOOS=linux GOARCH=amd64 go build -o binaries/amd64/linux/$${applicationName}-$${BUILD_NUMBER}.linux.amd64 $${DEPLOY_FILENAME} \
		&& echo "Building the Mac x64 binary" \
		&& GOOS=darwin GOARCH=amd64 go build -o binaries/amd64/darwin/$${applicationName}-$${BUILD_NUMBER}.darwin.amd64 $${DEPLOY_FILENAME} \
		&& echo "Building the Windows x64 binary" \
		&& GOOS=windows GOARCH=amd64 go build -o binaries/amd64/windows/$${applicationName}-$${BUILD_NUMBER}.windows.amd64.exe $${DEPLOY_FILENAME} \
		&& echo "Generating sha256sum" \
		&& sha256sum binaries/*/*/* > binaries/sha256sum.txt \
		&& echo "Done."
#		&& echo "Packaging static files" \
#		&& /go/bin/packr --verbose --input $${DEPLOY_FILENAME} \
#		&& /go/bin/packr clean
