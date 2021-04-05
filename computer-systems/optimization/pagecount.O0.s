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
	mov	qword ptr [rbp - 8], rdi
	mov	qword ptr [rbp - 16], rsi
	mov	rax, qword ptr [rbp - 8]
	xor	ecx, ecx
	mov	edx, ecx
	div	qword ptr [rbp - 16]
	pop	rbp
	ret
	.cfi_endproc
                                        ## -- End function
	.section	__TEXT,__literal8,8byte_literals
	.p2align	3               ## -- Begin function main
LCPI1_0:
	.quad	4711630319722168320     ## double 1.0E+7
LCPI1_1:
	.quad	4741671816366391296     ## double 1.0E+9
LCPI1_2:
	.quad	4696837146684686336     ## double 1.0E+6
	.section	__TEXT,__literal16,16byte_literals
	.p2align	4
LCPI1_3:
	.long	1127219200              ## 0x43300000
	.long	1160773632              ## 0x45300000
	.long	0                       ## 0x0
	.long	0                       ## 0x0
LCPI1_4:
	.quad	4841369599423283200     ## double 4503599627370496
	.quad	4985484787499139072     ## double 1.9342813113834067E+25
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
	sub	rsp, 192
	mov	rax, qword ptr [rip + ___stack_chk_guard@GOTPCREL]
	mov	rax, qword ptr [rax]
	mov	qword ptr [rbp - 8], rax
	mov	dword ptr [rbp - 68], 0
	mov	dword ptr [rbp - 72], edi
	mov	qword ptr [rbp - 80], rsi
	mov	dword ptr [rbp - 152], 0
	mov	rax, qword ptr [rip + l___const.main.msizes]
	mov	qword ptr [rbp - 32], rax
	mov	rax, qword ptr [rip + l___const.main.msizes+8]
	mov	qword ptr [rbp - 24], rax
	mov	rax, qword ptr [rip + l___const.main.msizes+16]
	mov	qword ptr [rbp - 16], rax
	mov	rax, qword ptr [rip + l___const.main.psizes]
	mov	qword ptr [rbp - 64], rax
	mov	rax, qword ptr [rip + l___const.main.psizes+8]
	mov	qword ptr [rbp - 56], rax
	mov	rax, qword ptr [rip + l___const.main.psizes+16]
	mov	qword ptr [rbp - 48], rax
	call	_clock
	mov	qword ptr [rbp - 88], rax
	mov	dword ptr [rbp - 148], 0
LBB1_1:                                 ## =>This Inner Loop Header: Depth=1
	cmp	dword ptr [rbp - 148], 10000000
	jge	LBB1_4
## %bb.2:                               ##   in Loop: Header=BB1_1 Depth=1
	mov	eax, dword ptr [rbp - 148]
	cdq
	mov	ecx, 3
	idiv	ecx
	movsxd	rsi, edx
	mov	rsi, qword ptr [rbp + 8*rsi - 32]
	mov	qword ptr [rbp - 120], rsi
	mov	edx, dword ptr [rbp - 148]
	mov	eax, edx
	cdq
	idiv	ecx
	movsxd	rsi, edx
	mov	rsi, qword ptr [rbp + 8*rsi - 64]
	mov	qword ptr [rbp - 128], rsi
	mov	rsi, qword ptr [rbp - 120]
	add	rsi, 1
	add	rsi, qword ptr [rbp - 128]
	movsxd	rdi, dword ptr [rbp - 152]
	add	rdi, rsi
                                        ## kill: def $edi killed $edi killed $rdi
	mov	dword ptr [rbp - 152], edi
## %bb.3:                               ##   in Loop: Header=BB1_1 Depth=1
	mov	eax, dword ptr [rbp - 148]
	add	eax, 1
	mov	dword ptr [rbp - 148], eax
	jmp	LBB1_1
LBB1_4:
	call	_clock
	mov	qword ptr [rbp - 96], rax
	call	_clock
	mov	qword ptr [rbp - 104], rax
	mov	dword ptr [rbp - 148], 0
LBB1_5:                                 ## =>This Inner Loop Header: Depth=1
	cmp	dword ptr [rbp - 148], 10000000
	jge	LBB1_8
## %bb.6:                               ##   in Loop: Header=BB1_5 Depth=1
	mov	eax, dword ptr [rbp - 148]
	cdq
	mov	ecx, 3
	idiv	ecx
	movsxd	rsi, edx
	mov	rsi, qword ptr [rbp + 8*rsi - 32]
	mov	qword ptr [rbp - 120], rsi
	mov	edx, dword ptr [rbp - 148]
	mov	eax, edx
	cdq
	idiv	ecx
	movsxd	rsi, edx
	mov	rsi, qword ptr [rbp + 8*rsi - 64]
	mov	qword ptr [rbp - 128], rsi
	mov	rdi, qword ptr [rbp - 120]
	mov	rsi, qword ptr [rbp - 128]
	call	_pagecount
	add	rax, qword ptr [rbp - 120]
	add	rax, qword ptr [rbp - 128]
	movsxd	rsi, dword ptr [rbp - 152]
	add	rsi, rax
                                        ## kill: def $esi killed $esi killed $rsi
	mov	dword ptr [rbp - 152], esi
## %bb.7:                               ##   in Loop: Header=BB1_5 Depth=1
	mov	eax, dword ptr [rbp - 148]
	add	eax, 1
	mov	dword ptr [rbp - 148], eax
	jmp	LBB1_5
LBB1_8:
	movsd	xmm0, qword ptr [rip + LCPI1_0] ## xmm0 = mem[0],zero
	movsd	xmm1, qword ptr [rip + LCPI1_1] ## xmm1 = mem[0],zero
	movsd	xmm2, qword ptr [rip + LCPI1_2] ## xmm2 = mem[0],zero
	movsd	qword ptr [rbp - 160], xmm0 ## 8-byte Spill
	movsd	qword ptr [rbp - 168], xmm1 ## 8-byte Spill
	movsd	qword ptr [rbp - 176], xmm2 ## 8-byte Spill
	call	_clock
	mov	qword ptr [rbp - 112], rax
	mov	rax, qword ptr [rbp - 112]
	mov	rcx, qword ptr [rbp - 104]
	sub	rax, rcx
	mov	rcx, qword ptr [rbp - 96]
	mov	rdx, qword ptr [rbp - 88]
	sub	rdx, rcx
	add	rax, rdx
	movq	xmm0, rax
	movaps	xmm1, xmmword ptr [rip + LCPI1_3] ## xmm1 = [1127219200,1160773632,0,0]
	punpckldq	xmm0, xmm1      ## xmm0 = xmm0[0],xmm1[0],xmm0[1],xmm1[1]
	movapd	xmm1, xmmword ptr [rip + LCPI1_4] ## xmm1 = [4.503599627370496E+15,1.9342813113834067E+25]
	subpd	xmm0, xmm1
	movaps	xmm1, xmm0
	unpckhpd	xmm0, xmm0      ## xmm0 = xmm0[1,1]
	addsd	xmm0, xmm1
	movsd	qword ptr [rbp - 136], xmm0
	movsd	xmm0, qword ptr [rbp - 136] ## xmm0 = mem[0],zero
	movsd	xmm1, qword ptr [rbp - 176] ## 8-byte Reload
                                        ## xmm1 = mem[0],zero
	divsd	xmm0, xmm1
	movsd	qword ptr [rbp - 144], xmm0
	movsd	xmm0, qword ptr [rbp - 144] ## xmm0 = mem[0],zero
	movsd	xmm2, qword ptr [rbp - 168] ## 8-byte Reload
                                        ## xmm2 = mem[0],zero
	mulsd	xmm2, qword ptr [rbp - 144]
	movsd	xmm3, qword ptr [rbp - 160] ## 8-byte Reload
                                        ## xmm3 = mem[0],zero
	divsd	xmm2, xmm3
	lea	rdi, [rip + L_.str]
	mov	esi, 10000000
	movaps	xmm1, xmm2
	mov	al, 2
	call	_printf
	mov	esi, dword ptr [rbp - 152]
	mov	rcx, qword ptr [rip + ___stack_chk_guard@GOTPCREL]
	mov	rcx, qword ptr [rcx]
	mov	rdx, qword ptr [rbp - 8]
	cmp	rcx, rdx
	mov	dword ptr [rbp - 180], esi ## 4-byte Spill
	jne	LBB1_10
## %bb.9:
	mov	eax, dword ptr [rbp - 180] ## 4-byte Reload
	add	rsp, 192
	pop	rbp
	ret
LBB1_10:
	call	___stack_chk_fail
	ud2
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
