// Definitions are very similar to structs and are primarily used for schemas.
// They are closed by default and are not emitted by Cue when exporting.

#schema: {
	word: string
	num: int | *42
	optional?: string
}

value: #schema & {
	word: "what's the good?"
}

// cue eval defns.cue
// cue export defns.cue