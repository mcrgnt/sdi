package sdi

type (
	Depser interface {
		Deps() []any
	}

	// Injector receives resolved dependencies from [Resolve]. Inject must only
	// store references — reading a dependency's own Inject-computed state here
	// is unsafe, since Resolve runs Inject across all resources in one pass
	// ordered by registration, not by the Deps() graph. See the lifecycle
	// safety rule in [github.com/omcrgnt/app]'s package doc.
	Injector interface {
		Inject([]any)
	}

	Compatible interface {
		Depser
		Injector
	}
)
