section .text
global index
index:
	; rdi: matrix
	; rsi: rows
	; rdx: cols
	; rcx: rindex
	; r8: cindex
	; formula is n_cols * r_idx + c_idx
	mov rax, rdx
	mul rcx
	add rax, r8
	; get it from memory, 4 bytes per increment
	mov rax, [rdi + 4*rax]
	ret
