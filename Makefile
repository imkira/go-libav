PACKAGES=$(shell find * -name *.go -print0 | xargs -0 -n1 dirname | sort --unique)
TEST_PACKAGES=$(shell find * -name *_test.go -print0 | xargs -0 -n1 dirname | sort --unique)

.PHONY: all gofmt golint govet test clean

all: gofmt golint govet test cover

fixtures:
	mkdir -p fixtures
	cd fixtures && rm -f tmp.zip && curl -o tmp.zip https://support.apple.com/library/APPLE/APPLECARE_ALLGEOS/HT1425/sample_iPod.m4v.zip && unzip -x tmp.zip && rm -f tmp.zip
	cd fixtures && rm -f tmp.zip && curl -o tmp.zip https://support.apple.com/library/APPLE/APPLECARE_ALLGEOS/HT1425/sample_iTunes.mov.zip && unzip -x tmp.zip && rm -f tmp.zip
	cd fixtures && rm -f tmp.zip && curl -o tmp.zip https://support.apple.com/library/APPLE/APPLECARE_ALLGEOS/HT1425/sample_mpeg4.mp4.zip && unzip -x tmp.zip && rm -f tmp.zip

gofmt:
	@for dir in $(PACKAGES); do gofmt -s=true -d=true -l=true $${dir}; done

golint:
	@for dir in $(PACKAGES); do golint $${dir}; done

govet:
	@for dir in $(PACKAGES); do go tool vet -all $${dir}; done

test: fixtures
	rm -f coverage.*
	@for dir in $(TEST_PACKAGES); do (cd $${dir} && go test -v -race -cpu=1,2,4 -coverprofile=coverage.txt -covermode=atomic || touch $(PWD)/coverage.failed); done
	@for dir in $(TEST_PACKAGES); do (if [ -f coverage.txt ]; then cat $${dir}/coverage.txt | tail -n +2 >> coverage.txt; else cp $${dir}/coverage.txt .; fi); done
	@test ! -f coverage.failed || (echo Tests failed; exit 1)

cover:
	go tool cover -html=coverage.txt -o coverage.html

clean:
	rm -rf fixtures
