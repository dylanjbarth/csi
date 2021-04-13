loop: loop-order.c
	gcc -g loop-order.c -o loop

grind: loop
	valgrind --tool=cachegrind ./loop

time: loop
	time ./loop

clean:
	rm loop

loop-order.c: 