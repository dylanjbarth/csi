PROGRAM=ls-clone.out

clean: 
	rm -rf $(PROGRAM)*

build: clean
	gcc -g -o $(PROGRAM) ls-clone.c

run: build
	./$(PROGRAM) . .. ls-clone.c notfile cantaccess

debug: build
	lldb ./$(PROGRAM)