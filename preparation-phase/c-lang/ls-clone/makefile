PROGRAM=ls-clone.out

clean: 
	rm -rf $(PROGRAM)*

build: clean
	gcc -g main.c sorting.c -o $(PROGRAM) 

test: build
	./test.sh

debug: build
	lldb ./$(PROGRAM)
