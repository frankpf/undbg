global _start

section .text

exit:
	mov rax, 60
	mov rdi, 1
	syscall

_start:
	xor rax, rax
	add rax, 10
	sub rax, 9
	mov rdi, 1
	mov rsi, hello_world
	mov rdx, length
	syscall
	jmp exit

section .data
	hello_world: db 'hello world', 0xa
	length: equ $-hello_world
