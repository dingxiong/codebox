# Pratt Parser for C expressions

This code is mostly based on Alex's Rust code 
example https://matklad.github.io/2020/04/13/simple-but-powerful-pratt-parsing.html 
with some differences:
1. Python's enum is not as powerful as Rust's enum. I instead use class hiarchery.
2. I do not have `EOF` token. For an empty input, the parser return `None`.

C's precedence table is [here](https://en.cppreference.com/w/c/language/operator_precedence). 
So far I only finished a subset of the implementation.

One problem with Pratt parser is that 

