section .text
global fib
fib:
	; base case -- if rax is <= 1, return  
	cmp rdi, 1  
	jnbe .recurse
	mov rax, rdi
	ret
.recurse: ; return fib(n-1) + fib(n-2)
	push rdi ; copy original rdi onto stack before we modify it
	sub rdi, 1 ; fib(n-1)
	call fib
	pop rdi ; pop our original rdi off the stack before next call
	push rax  ; save the result to the stack before it's overwritten by the next call
	sub rdi, 2  ; fib(n-2)
	call fib
	pop r8  ; pop fib(n-1) off stack for addition
	add rax, r8  ; add them together
	ret
