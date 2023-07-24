// Cue has list comprehensions to dynamically create lists. You can iterate over both lists and struct fields.
//
//	The form is [ for key, val in <iterable> [condition] { production } ]
//
//	* key is the index for lists and the label for fields

nums: [1, 2, 3, 4, 5, 6]
sqrd: [ for _, n in nums{n * n}]
even: [ for _, n in nums if mod(n, 2) == 0 {n}]

listOfStructs: [
	for p, n in nums {
		position: p
		value: n
	}
]

extractVals: [ for p, S in listOfStructs {S.value}]