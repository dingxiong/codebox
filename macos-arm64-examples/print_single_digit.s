// Every basic program to print out a single digit.
// It does not use glibc functions.
// 
// how to run 
// as -o print_single_digit.o print_single_digit.s  
// ld -o print_single_digit print_single_digit.o -lSystem -syslibroot `xcrun -sdk macosx --show-sdk-path` -e _start -arch arm64

.global _start
.p2align 3 

.data 
my_number: .byte 5

.text
_start:
    adrp    x0, my_number@PAGE      // Get high bits of address
    add     x0, x0, my_number@PAGEOFF // Add low bits of offset
    
    ldr x0, [x0]

    add x0, x0, #48   // Convert the digit to ASCII. '0' = 48.
    str x0, [sp, #-16]! 

    // Print the digit
    mov x0, #1        // File descriptor (stdout)
    mov x1, sp     // Buffer address
    mov x2, #1        // Length
    mov x16, #4       // write syscall number
    svc #0            // Make syscall

    // Exit program
    mov x0, #0        // Exit code
    mov x16, #1       // exit syscall number
    svc #0            // Call MacOS to terminate the program

