.PHONY: fmt clean contributors

BUILDTAGS=debug

fmt:
	go fmt amald/...

clean:
	go clean -i -r amald/...

deps:
	go get -tags '$(BUILDTAGS)' -d -v amald/... 

server: deps
	go install -tags '$(BUILDTAGS)' amald/main/amald

test:
	go test ./src/amald/...

contributors:
	@echo "Amald has been created by the following fine individuals:\n" > CONTRIBUTORS
	git log --raw | grep "^Author: " | sort | uniq | cut -d ' ' -f2- | sed 's/^/- /' | cut -d '<' -f1 >> CONTRIBUTORS
