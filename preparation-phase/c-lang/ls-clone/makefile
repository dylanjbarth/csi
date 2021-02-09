PROGRAM=ls-clone.out

clean: 
	rm -rf $(PROGRAM)*

build: clean
	gcc -g main.c sorting.c -o $(PROGRAM) 

test: build
	./test.sh

# test: build build-tests
# 	./$(PROGRAM) .
# 	./$(PROGRAM) test
# 	./$(PROGRAM) -Ca test
# 	./$(PROGRAM) -C test
# 	./$(PROGRAM) -1 test
# 	./$(PROGRAM) -1a test
# 	./$(PROGRAM) -f test
# 	./$(PROGRAM) -f1 test
# 	./$(PROGRAM) -S1a test
# 	./$(PROGRAM) -Sa test

debug: build
	lldb ./$(PROGRAM)
