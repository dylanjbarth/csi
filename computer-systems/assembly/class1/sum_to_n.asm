section .text
global sum_to_n
sum_to_n:
	; Set counter to 0 -- is this needed or does it always start here? seems like it fails if not set
	mov rcx, 0
	; mov rax, 0
.add_v:
	cmp rcx, rdi
	jg .done
	add rax, rcx
	inc rcx
	jmp .add_v
.done:
	ret
