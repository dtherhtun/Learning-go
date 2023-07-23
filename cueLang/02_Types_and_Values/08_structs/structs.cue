// Structs are like JSON objects. They are the primary composite type in Cue.
// They have a set of fields (label: value). By default, they are open and you can add more fields.

// an open struct
a: {
	foo: "bar"
}

// shorthand nested field
a: hello: "world"

// a closed struct
b: close({
	left: "right"
})

// error, field up not allowed
b: up: "down"