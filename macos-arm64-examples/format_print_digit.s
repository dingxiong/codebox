// 
// how to run 
// as -o format_print_digit.o format_print_digit.s
// ld -o format_print_digit format_print_digit.o -lSystem -syslibroot `xcrun -sdk macosx --show-sdk-path` -e _start -arch arm64

.global _start
.p2align 3 
.extern _printf
.extern _exit

.data
my_str: .asciz "I am printing a digit: %d\n"
my_byte: .byte 3

.text
_start: 
    // Maintain 16-byte stack alignment for ABI compliance
    stp x29, x30, [sp, #-16]!   // Pre-decrement SP and store FP and LR
    mov x29, sp                  // Set up frame pointer

    // Set up printf arguments
    adrp x0, my_str@PAGE        // Load high bits of string address
    add x0, x0, my_str@PAGEOFF  // Complete the address

    adrp x1, my_byte@PAGE        // Load high bits of string address
    add x1, x1, my_byte@PAGEOFF  // Complete the address
    ldr x1, [x1]                  

    // According to MacOS ARM64 ABI 
    // https://stackoverflow.com/questions/69454175/calling-printf-from-aarch64-asm-code-on-apple-m1-macos
    // variadic arguments need to be passed to stack as well
    str x1, [sp, #-16]!
    bl _printf
    add sp, sp, #16

    // Clean up and exit
    mov w0, #0                  // Exit status 0
    ldp x29, x30, [sp], #16    // Restore FP and LR, post-increment SP
    bl _exit
