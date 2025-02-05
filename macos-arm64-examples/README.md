
Some example assembly code for Mac AArch64 chip.

## For loop
For the simple loop `for(init; condition; action) { body }`, the assembly pseudocode is 
```
init
jmp L2
L1:
    action 
L2: 
    condtion;
    jmp L3 if false
    body 
    jmp L1
L3:
```
