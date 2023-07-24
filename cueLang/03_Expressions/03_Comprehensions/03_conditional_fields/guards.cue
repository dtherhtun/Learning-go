// CUE’s if is different from other languages.
// It is a comprehension rather than a branching mechanism.
// That is why we refer to it as conditional fields, guarded fields,
// or another form of field comprehension.
//
// Some important differences:
//
// - there is no else statement, you only include config when statements are true
// - there is no short-circuiting for multiple checks, all conditions will be evaluated

app: {
	name: string
	tech: string
	mem: int

	if tech == "react" {
		tier: "frontend"
	}
	if tech != "react" {
		tier: "backend"
	}

	if mem < 1Gi {
		footprint: "small"
	}
	if mem >= 1Gi && mem <= 4Gi {
		footprint: "medium"
	}
	if mem >= 4Gi {
		footprint: "large"
	}
}

// This will result in an error because CUE evaluates all conditions
// without short-circuiting, meaning it will still try to access app.field
// if app.field != _|_ && app.field == true {
//   foo: true
// }

// Use nested guards to check multiple conditions
if app.field != _|_ {
	if app.field == true {
		foo: true
	}
}