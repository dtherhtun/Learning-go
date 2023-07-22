// Definitions are Cue’s way of specifying schemas.
// They have slightly different rules from structs.
//
// - They are not output as data
// - They may remain incomplete or under specified
// - 	They “close” a struct, forbidding unknown or additional fields

#Album: {
	artist: string
	title: string
	year: int
}

album: #Album & {
	artist: "Erutan"
	title: "Always With Me (Itsumo Nando Demo)"
	year: 2001
}

// cue export 05_definitions/definition.cue --out json