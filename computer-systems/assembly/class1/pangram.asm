; starting with most straightforward solution, flipping bits for the 26 characters

section .text
global pangram
pangram:
	; strategy is to iterate through the sentence, flipping bits in memory and then seeing if all are set
	mov rax, 0 ; zero out our bit flipper
.checkletter:
	movzx ecx, byte [rdi]
	; if null, we are ready to check the total
	cmp ecx, 0
	je .total

	; otherwise check conditions that mean we can just skip to next
	or ecx, 32  ; lowercase
	cmp ecx, 'a'
	jl .loopcontinue
	cmp ecx, 'z'
	jg .loopcontinue

.isupper: ; at this point we've trimmed everything out that isn't A-Z
	; mark the bits in memory 
	sub ecx, 'a' ; get the letter index from 0
	bts eax, ecx  ; set the bit to 1 at the index of ecx
.loopcontinue:
	inc rdi
	jmp .checkletter
.total:
	; true if letters are all 1 otherwise false
	cmp rax, 0x3ffffff  ; 2^26
	je .true
	; else 
	mov rax, 0
	jmp .done
.true:
	mov rax, 1
	jmp .done
.done:
	ret

