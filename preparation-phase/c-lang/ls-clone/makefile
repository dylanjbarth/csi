PROGRAM=ls-clone.out

clean: 
	rm -rf $(PROGRAM)*

clean-tests: 
	rm -rf test

build: clean
	gcc -g main.c sorting.c -o $(PROGRAM) 

test: build build-tests
	./$(PROGRAM) .
	./$(PROGRAM) test
	./$(PROGRAM) -Ca test
	./$(PROGRAM) -C test
	./$(PROGRAM) -1 test
	./$(PROGRAM) -1a test
	./$(PROGRAM) -f test
	./$(PROGRAM) -f1 test
	./$(PROGRAM) -S1a test
	./$(PROGRAM) -Sa test

debug: build
	lldb ./$(PROGRAM)

build-tests: clean-tests
	mkdir test
	# Create a bunch of files.
	nfiles=1 ; while [[ $$nfiles -le 5 ]]; do \
		echo "Test $$nfiles" > test/testfile$$nfiles.txt; \
		((nfiles = nfiles + 1)) ; \
	done
	# create a few more with interesting permissions or file lengths.
	touch test/areallllllllllllllllllllllllllllllylongfilename.txt && chmod 0777 test/areallllllllllllllllllllllllllllllylongfilename.txt
	touch test/global_all.txt && chmod 0777 test/global_all.txt
	touch test/user_read.txt && chmod 0400 test/user_read.txt
	touch test/group_read.txt && chmod 0740 test/group_read.txt
	echo "A bit biggggggggggggggggger" > test/bigfile.txt
	echo "A bit smaller" > test/smallfile.txt
	# test another directory
	mkdir test/another_dir
