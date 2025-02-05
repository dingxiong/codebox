// 
// how to run 
// as -o loop.o loop.s
// ld -macos_version_min 15.0.0 -lSystem -syslibroot `xcrun -sdk macosx --show-sdk-path` -e _start -arch arm64 -o loop loop.o

// for (int i = 0; i < 5; i++) { printf("%d\n", i); }
.global _start
.p2align 3 
.extern _printf
.extern _exit

.data
my_str: .asciz "%d\n"

.text
_start: 
    sub sp, sp, #-32
    stp x29, x30, [sp, #-16] 
    add x29, sp, #16 

    // int i = 0;
    mov x0, #0
    str x0, [x29, #-8]

    b L2 
L1:
    // i++
    ldr x0, [x29, #-8]
    add x0, x0, #1 
    str x0, [x29, #-8]
L2:
    // i < 5 
    ldr x0, [x29, #-8]
    cmp x0, #5
    BGE L3

    // print("%d\n", i)
    adrp x0, my_str@PAGE        
    add x0, x0, my_str@PAGEOFF 

    ldr x1, [x29, #-8]
    
    str x1, [sp, #-16]!
    bl _printf
    add sp, sp, #16
    
    b L1
L3:

    // Clean up and exit
    mov w0, #0              
    ldp x29, x30, [sp, #16] 
    add sp, sp, #32
    bl _exit
