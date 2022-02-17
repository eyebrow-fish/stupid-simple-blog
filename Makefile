all: clean build build-test run-test

clean:
	rm -rf out/

build:
	mkdir out/
	mkdir out/target/
	cc -Wall -o out/target/stupid-simple-blog src/*.c -I src/

build-test: build
	mkdir out/test/
	cc -Wall -o out/test/stupid-simple-blog src/httprequest.c src/httpresponse.c test/*.c -I src/ -I test/

run-test:
	./out/test/stupid-simple-blog

run:
	./out/target/stupid-simple-blog
