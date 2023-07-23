// Open means a struct can be extended, closed means they cannot.
// By default, structs are open and definitions are closed.
// Cue also allows us to explicitly do the opposite.

// Closed struct
s: close({
	foo: "bar"
})

// Open definition
#d: {
	foo: "bar"
	... // must be last
}