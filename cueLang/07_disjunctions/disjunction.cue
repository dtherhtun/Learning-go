// disjunction of values (like an enum)
hello: "world" | "bob" | "mary"
hello: "world"

// disjunction of types
port: string | int
port: 5432

// disjunction of schemas
val: #Def1 | #Def2
val: {foo: "bar", ans: 42}

#Def1: {
	foo: string
	ans: int
}

#Def2: {
	name: string
	port: int
}

// Disjunctions have several uses:
//
// - enums (as values)
// - sum-type (any of these types)
// - null-coalescing (use this computation, or default to some value)