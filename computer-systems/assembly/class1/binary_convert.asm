; naive approach first: get length of string, then work backward incrementing by powers of 2
; better approach, bit shift left to raise to power of two

section .text
global binary_convert
binary_convert:	
	; get string length
	mov r9, 0  ; displacement (and will hold length of input)
	mov r10d, 1  ; scalar
	mov r11d, 0  ; sum
.strlen:
	movzx ecx, byte [rdi + r9]
	cmp ecx, 0  ; check if null
	je .sum
	inc r9
	jmp .strlen
.sum:
	; populate eax with a 1 or 0 depending on the string character byte 
	movzx ecx, byte [rdi + r9 - 1]
	cmp ecx, 48
	je .casezero
	jmp .caseone
.casezero:
	mov eax, 0
	jmp .scale
.caseone:
	mov eax, 1
	; multiply by our scalar (based on the place value) and cache result in r11d
.scale:
	mul r10d	
	add r11d, eax
	; increment our scalar by the next place
	mov eax, r10d
	imul eax, 2
	mov r10d, eax
	; move to the next place
	dec r9
	cmp r9, 0  ; termination condition
	jne .sum
	; when we've added everything move our sum to the output register
	mov eax, r11d	
	ret 

section .data
zero: db 48

