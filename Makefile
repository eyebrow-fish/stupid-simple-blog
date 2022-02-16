full-build: clean clean-test build build-test run-test

clean:
	rm -rf out/target/

clean-test:
	rm -rf out/test/

build:
	mkdir out/target/
	cc -Wall -o out/target/stupid-simple-blog src/*.c -I src/

build-test:
	mkdir out/test/
	cc -Wall -o out/test/stupid-simple-blog test/*.c -I src/ -I test/

run-test:
	./out/test/stupid-simple-blog

run:
	./out/target/stupid-simple-blog
