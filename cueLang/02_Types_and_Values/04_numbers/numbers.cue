// Cue defines two number kinds:
//
// - int are whole numbers, implemented as BigInt to represent any value,
//   and can be constrained for byte size (like int64)
// - number are decimals numbers, (also not bounded by byte size, also constrainable?),
//   ints are also numbers

a: int
a: 42

b: number
b: 3.14

c: int & 42.0 // conflict int and 42.0

d: 42 // will be type int
e: 3.14 // will be type number
