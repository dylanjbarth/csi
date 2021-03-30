; starting with most straightforward solution, array of 26 bytes (because I'm not sure how to address individual bits yet...)

section .text
global pangram
pangram:
	; strategy is to iterate through the sentence, flipping bits in memory and then seeing if all are set
	mov r8, 0  ; counter
.checkletter:
	movzx ecx, byte [rdi + r8]
	; if null, we are done
	cmp ecx, 0
	je .done

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
	; mark the bytes in memory 
	sub ecx, 65 ; get the letter index from 0
	mov byte [letters], 1  ; flip the bit
.loopcontinue:
	inc r8
	jmp .checkletter
.done:
	; true if letters are all 1 otherwise false
	ret

; section .data
; A: 65
; a: 97
; Z: 90
; z: 122

section .bss
letters: resb 26
