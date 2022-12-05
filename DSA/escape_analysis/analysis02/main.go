package main

// returnValue returns a value over the call stack
func returnValue() int {
    n := 42
    return n
}

// returnPointer returns a pointer to n, and n is moved to the heap
func returnPointer() *int {
    n := 42  // --- line 11 ---
    return &n
}

// returnSlice returns a slice that escapes to the heap,
// because the slice header includes a pointer to the data.
func returnSlice() []int {
    slice := []int{42} // --- line 18 ---
    return slice
}

// returnArray returns an array that does not escape to the heap,
// because arrays need no header for tracking length and capacity
// and are always copied by value.
func returnArray() [1]int {
    return [1]int{42}
}

// largeArray creates a ridiculously large array that escapes to the heap,
// even though the array itself is not returned
// and thus does not outlive the function.
func largeArray() int {
    var largeArray [100000000]int  // --- line 33 ---
    largeArray[42] = 42
    return largeArray[42]
}

// returnFunc() returns a function that escapes to the heap
func returnFunc() func() int {
    f := func() int {  // --- line 40 ---
        return 42
    }
    return f
}

func main() {
    a := returnValue()
    p := *returnPointer()
    s := returnSlice()
    arr := returnArray()
    la := largeArray()
    f := returnFunc()()

    // Consume the variables to avoid compiler warnings.
    // I don't use Printf/ln because this produces a lot 
    // of extra escape messages.
    if a+p+s[0]+arr[0]+la+f == 0 {
        return
    }
}
