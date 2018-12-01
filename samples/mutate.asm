global _start

section .text

exit:
	mov rax, 60
	mov rdi, 20
	syscall

info:
	mov rax, 1
	mov rdi, 0
	mov rsi, test_str
	mov rdx, test_str_len
	syscall
	ret

_start:
	xor rax, rax
	lea rax, [mem_byte]
	mov qword [rax], 1234
	cmp qword [rax], 4321
	jne else
	call info
else:
	mov qword [rax], 4321
	jmp exit

section .data
	test_str: db 'hello world', 0xa
	test_str_len: equ $-test_str

section .bss
	mem_byte: resb 4
