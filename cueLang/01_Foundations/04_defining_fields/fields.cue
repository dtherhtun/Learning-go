// - basic data types must be the same
// - you can make a field more restrictive, but not the other way
// - structs fields are merged, list elements must match exactly
// - the rules are applied recursively

hello: "world"
hello: "world"

// set a type
s: {a: int}

// set some data
s: {a: 1, b: 2}

// set a nested field without curly braces
s: c: d: 3

// list must have the same elements
// and cannot change length
l: ["abc", "123"]
l: [
	"abc",
	"123"
]