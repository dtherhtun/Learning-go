// Cue has the following built in types.
// They are meant to align with JSON types.
//
// null  bool  string  bytes  number  list  struct
//                              |
//                             int

N: null
B: bool
S: string
By: bytes
Num: number // Decimals or integers, a superclass if you will
Int: int    // Big Int which can represent values without limits
List: [...]
Struct: {...}