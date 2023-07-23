// An incomplete value is one which does not have all fields filled with data.
// Cue will not export incomplete values and instead return an error.
// By contrast, concrete is a fully specified value.

// incomplete value
a: _
b: int

s: {
	a: _
}

// concreate value
a: "a"
b: int

s: a: {foo: "bar"}