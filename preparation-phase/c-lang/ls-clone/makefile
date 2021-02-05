PROGRAM=ls-clone.out

clean: 
	rm -rf $(PROGRAM)*

build: clean
	gcc -g -o $(PROGRAM) ls-clone.c

run: build
	./$(PROGRAM) 1 2 3

debug: build
	lldb ./$(PROGRAM)