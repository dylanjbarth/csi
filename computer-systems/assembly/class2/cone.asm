default rel

section .text
global volume
volume:
	; a function that calculates the volume of a cone, given its radius and height as the first two arguments.
	; The first and second arguments are provided in xmm0 and xmm1 respectively. xmm0 is reused for the return value.
	; V=π * r^2 * h/3
	mulss xmm0, xmm0  ; r^2
	mulss xmm0, [pi] ; * pi
	mulss xmm0, xmm1  ; * h
	divss xmm0, [three] ; /3
 	ret

section .rodata
pi: dd 3.14159
three: dd 3.0
