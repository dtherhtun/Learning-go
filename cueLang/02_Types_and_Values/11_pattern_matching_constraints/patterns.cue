// Pattern matching allows you to specify constraints for labels which match a pattern.

#Schema: {
	name: string
	ans: string
	num: int | *42
}

// match elem fields and alias labels to Name,
// unify with schema, set name to Name by label
elems: [Name=_]: #Schema & {name: Name}

elems: {
	one: {
		ans: "solo"
		num: 1
	}
	two: {
		ans: "life"
	}
}

elems: other: {ans: "id", num: 23}