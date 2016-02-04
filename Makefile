PACKAGES=$(shell find * -name *.go -print0 | xargs -0 -n1 dirname | sort --unique)
TEST_PACKAGES=$(shell find * -name *_test.go -print0 | xargs -0 -n1 dirname | sort --unique)

FIXTURES=sample_iPod.m4v        \
         sample_iTunes.mov      \
         sample_mpeg4.mp4

.PHONY: all gofmt golint govet test clean

all: gofmt golint govet test cover

gofmt:
	@for dir in $(PACKAGES); do gofmt -s=true -d=true -l=true $${dir}; done

golint:
	@for dir in $(PACKAGES); do golint $${dir}; done

govet:
	@for dir in $(PACKAGES); do go tool vet -all $${dir}; done

FIXTURE_TARGETS=$(addprefix fixtures/,$(FIXTURES))

$(FIXTURE_TARGETS):
	mkdir -p "$(dir $@)"
	rm -f "$@.zip" "$@"
	cd "$(dir $@)" && curl -O "https://support.apple.com/library/APPLE/APPLECARE_ALLGEOS/HT1425/$(notdir $@).zip" && unzip $(notdir $@).zip
	rm -f "$@.zip"

fixtures: $(FIXTURE_TARGETS)

test: fixtures
	rm -f coverage.*
	@for dir in $(TEST_PACKAGES); do (cd $${dir} && go test -v -race -cpu=1,2,4 -coverprofile=coverage.txt -covermode=atomic || touch $(PWD)/coverage.failed); done
	@for dir in $(TEST_PACKAGES); do (if [ -f coverage.txt ]; then cat $${dir}/coverage.txt | tail -n +2 >> coverage.txt; else cp $${dir}/coverage.txt .; fi); done
	@test ! -f coverage.failed || (echo Tests failed; exit 1)

cover:
	go tool cover -html=coverage.txt -o coverage.html

clean:
	rm -rf fixtures
