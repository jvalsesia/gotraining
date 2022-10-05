# Pointers
Pointers provide a way to share data across program boundaries. Having the ability to share and reference data with a pointer provides the benefit of efficiency. There is only one copy of the data and everyone can see it changing. The cost is that anyone can change the data which can cause side effects in running programs.

## Notes
- Use pointers to share data.
- Values in Go are always pass by value.
- "Value of", what's in the box. "Address of" ( & ), where is the box.
- The (*) operator declares a pointer variable and the "Value that the pointer points to


## Exercises
### Exercise 1
**Part A** Declare and initialize a variable of type int with the value of 20. Display the address of and value of the variable.

**Part B** Declare and initialize a pointer variable of type int that points to the last variable you just created. Display the address of , value of and the value that the pointer points to.

### Exercise 2
Declare a struct type and create a value of this type. Declare a function that can change the value of some field in this struct type. Display the value before and after the call to your function.
