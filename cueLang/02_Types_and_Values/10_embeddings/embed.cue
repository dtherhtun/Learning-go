// You can embed structs and definitions within each other as a method to build up values.
// You can achieve the same with opened structs / definitions and conjunctions,
// but often we cannot modify what we can embed.
#A: {
	num: number
}

#B: {
	ans: string
}

// this won't work
//#bad: #A & #B
//bad:  #bad & {
//	num: 42
//	ans: "life"
//}

// but this will
#val: {#A, #B}
val: #val & {
	num: 42
	ans: "life"
}