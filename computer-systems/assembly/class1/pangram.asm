; starting with most straightforward solution, flipping bits for the 26 characters

section .text
global pangram
pangram:
	; strategy is to iterate through the sentence, flipping bits in memory and then seeing if all are set
	mov r8, 0  ; counter
	mov rax, 0 ; zero out our bit flipper
.checkletter:
	movzx ecx, byte [rdi + r8]
	; if null, we are ready to check the total
	cmp ecx, 0
	je .total

	; otherwise check conditions that mean we can just skip to next
	cmp ecx, 122
	jg .loopcontinue
	cmp ecx, 65
	jl .loopcontinue
	; now if greater than 90, we can subtract 32 to uppercase it.. then trim off the end. 
	cmp ecx, 90
	jle .isupper
	sub ecx, 32
	cmp ecx, 90
	jg .loopcontinue
	cmp ecx, 65
	jl .loopcontinue

.isupper: ; at this point we've trimmed everything out that isn't A-Z
	; mark the bits in memory 
	sub ecx, 65 ; get the letter index from 0
	bts eax, ecx  ; set the bit to 1 at the index of ecx
.loopcontinue:
	inc r8
	jmp .checkletter
.total:
	; true if letters are all 1 otherwise false
	cmp rax, 0x0000000003ffffff  ; 2^26
	je .true
	; else 
	mov rax, 0
	jmp .done
.true:
	mov rax, 1
	jmp .done
.done:
	ret

; section .data
; A: 65
; a: 97
; Z: 90
; z: 122
