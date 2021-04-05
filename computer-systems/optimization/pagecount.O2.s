	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 15	sdk_version 10, 15, 6
	.intel_syntax noprefix
	.globl	_pagecount              ## -- Begin function pagecount
	.p2align	4, 0x90
_pagecount:                             ## @pagecount
	.cfi_startproc
## %bb.0:
	push	rbp
	.cfi_def_cfa_offset 16
	.cfi_offset rbp, -16
	mov	rbp, rsp
	.cfi_def_cfa_register rbp
	mov	rax, rdi
	xor	edx, edx
	div	rsi
	pop	rbp
	ret
	.cfi_endproc
                                        ## -- End function
	.section	__TEXT,__literal16,16byte_literals
	.p2align	4               ## -- Begin function main
LCPI1_0:
	.long	1127219200              ## 0x43300000
	.long	1160773632              ## 0x45300000
	.long	0                       ## 0x0
	.long	0                       ## 0x0
LCPI1_1:
	.quad	4841369599423283200     ## double 4503599627370496
	.quad	4985484787499139072     ## double 1.9342813113834067E+25
	.section	__TEXT,__literal8,8byte_literals
	.p2align	3
LCPI1_2:
	.quad	4696837146684686336     ## double 1.0E+6
LCPI1_3:
	.quad	4741671816366391296     ## double 1.0E+9
LCPI1_4:
	.quad	4711630319722168320     ## double 1.0E+7
	.section	__TEXT,__text,regular,pure_instructions
	.globl	_main
	.p2align	4, 0x90
_main:                                  ## @main
	.cfi_startproc
## %bb.0:
	push	rbp
	.cfi_def_cfa_offset 16
	.cfi_offset rbp, -16
	mov	rbp, rsp
	.cfi_def_cfa_register rbp
	push	r15
	push	r14
	push	r13
	push	r12
	push	rbx
	sub	rsp, 24
	.cfi_offset rbx, -56
	.cfi_offset r12, -48
	.cfi_offset r13, -40
	.cfi_offset r14, -32
	.cfi_offset r15, -24
	xor	r12d, r12d
	mov	r14d, 8
	mov	ebx, 1
	mov	r13d, 10000000
	call	_clock
	mov	qword ptr [rbp - 56], rax ## 8-byte Spill
	movabs	r8, -6148914691236517205
	lea	r9, [rip + l___const.main.msizes]
	lea	rdi, [rip + l___const.main.psizes]
	xor	ecx, ecx
	xor	r15d, r15d
	.p2align	4, 0x90
LBB1_1:                                 ## =>This Inner Loop Header: Depth=1
	mov	rax, rbx
	mul	r8
	shl	rdx, 2
	and	rdx, -8
	lea	rax, [rdx + 2*rdx]
	mov	rsi, r14
	sub	rsi, rax
	mov	rax, r12
	mul	r8
	shl	rdx, 2
	and	rdx, -8
	lea	rax, [rdx + 2*rdx]
	mov	rdx, rcx
	sub	rdx, rax
	mov	eax, dword ptr [r9 + rdx]
	add	eax, dword ptr [rdi + rdx]
	add	eax, r15d
	mov	edx, dword ptr [r9 + rsi]
	add	edx, dword ptr [rdi + rsi]
	lea	r15d, [rax + rdx + 2]
	add	r14, 16
	add	rbx, 2
	add	rcx, 16
	add	r12, 2
	add	r13d, -2
	jne	LBB1_1
## %bb.2:
	mov	r13, r9
	mov	r12, rdi
	call	_clock
	mov	qword ptr [rbp - 48], rax ## 8-byte Spill
	mov	r14d, 10000000
	xor	ebx, ebx
	call	_clock
	mov	r10, r12
	mov	r9, r13
	movabs	r8, -6148914691236517205
	mov	r12, rax
	xor	esi, esi
	.p2align	4, 0x90
LBB1_3:                                 ## =>This Inner Loop Header: Depth=1
	mov	rax, rbx
	mul	r8
	shl	rdx, 2
	and	rdx, -8
	lea	rax, [rdx + 2*rdx]
	mov	rdx, rsi
	sub	rdx, rax
	mov	rcx, qword ptr [r9 + rdx]
	mov	rdi, qword ptr [r10 + rdx]
	mov	rax, rcx
	xor	edx, edx
	div	rdi
	add	edi, ecx
	add	eax, edi
	add	r15d, eax
	add	rsi, 8
	inc	rbx
	dec	r14d
	jne	LBB1_3
## %bb.4:
	call	_clock
	mov	rcx, qword ptr [rbp - 56] ## 8-byte Reload
	sub	rcx, qword ptr [rbp - 48] ## 8-byte Folded Reload
	sub	rcx, r12
	add	rcx, rax
	movq	xmm1, rcx
	punpckldq	xmm1, xmmword ptr [rip + LCPI1_0] ## xmm1 = xmm1[0],mem[0],xmm1[1],mem[1]
	subpd	xmm1, xmmword ptr [rip + LCPI1_1]
	movapd	xmm0, xmm1
	unpckhpd	xmm0, xmm1      ## xmm0 = xmm0[1],xmm1[1]
	addsd	xmm0, xmm1
	divsd	xmm0, qword ptr [rip + LCPI1_2]
	movsd	xmm1, qword ptr [rip + LCPI1_3] ## xmm1 = mem[0],zero
	mulsd	xmm1, xmm0
	divsd	xmm1, qword ptr [rip + LCPI1_4]
	lea	rdi, [rip + L_.str]
	mov	esi, 10000000
	mov	al, 2
	call	_printf
	mov	eax, r15d
	add	rsp, 24
	pop	rbx
	pop	r12
	pop	r13
	pop	r14
	pop	r15
	pop	rbp
	ret
	.cfi_endproc
                                        ## -- End function
	.section	__TEXT,__const
	.p2align	4               ## @__const.main.msizes
l___const.main.msizes:
	.quad	4294967296              ## 0x100000000
	.quad	1099511627776           ## 0x10000000000
	.quad	4503599627370496        ## 0x10000000000000

	.p2align	4               ## @__const.main.psizes
l___const.main.psizes:
	.quad	4096                    ## 0x1000
	.quad	65536                   ## 0x10000
	.quad	4294967296              ## 0x100000000

	.section	__TEXT,__cstring,cstring_literals
L_.str:                                 ## @.str
	.asciz	"%.2fs to run %d tests (%.2fns per test)\n"

.subsections_via_symbols
