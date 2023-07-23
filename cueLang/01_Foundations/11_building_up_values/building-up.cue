// In Cue, it is recommended to start small and build values up.
// This makes schemas reusable. You can do this by embedding values.

#Base: {
	name: string
	kind: string
}

#Meta: {
	// string and a semver regex
	version: string & =~"^v[0-9]+\\.[0-9]+\\.[0-9]+$"
	// list of strings
	labels: [...string]
}

#Permissions: {
	role: string
	public: bool | *false
}

#Schema: {
	// embed other schemas
	#Base
	#Meta

	#Permissions
	// with no '...' this is final
}

value: #Schema & {
	name: "app"
	kind: "deploy"
	version: "v1.0.42"
	labels: ["server", "prod"]
	role: "backend"
	// public false by default
}