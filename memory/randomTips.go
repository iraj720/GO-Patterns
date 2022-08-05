package memory

// Go compilers will allocate variables that are local to a function in that function’s stack frame. However,
// if the compiler cannot prove that the variable is not referenced after the function returns, then
// the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors

// A general rule we can infer from this is that sharing pointers up the stack results in allocations,
// whereas sharing points down the stack doesn’t. However, this is not guaranteed
