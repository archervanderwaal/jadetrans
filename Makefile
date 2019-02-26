GOCMD=go
BINARY=jadetrans
VERSION=V1.0
build:
	$(GOCMD) build
install:
	$(GOCMD) install
release:
	# clean *.tar.gz
	rm -rf *.gz
	# Build for mac
	$(GOCMD) clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
	tar czf $(BINARY)-mac64-$(VERSION).tar.gz $(BINARY)
	# Build for linux
	$(GOCMD) clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
	tar czf $(BINARY)-linux64-$(VERSION).tar.gz $(BINARY)
	# Build for windows
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
	tar czf $(BINARY)-win64-$(VERSION).tar.gz $(BINARY).exe
	$(GOCMD) clean
clean:
	$(GOCMD) clean
	rm -rf *.gz
