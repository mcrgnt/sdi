/*
Пакет sdi — wire-only DI: Inject; optional [CheckCycles] lint.

sdi не строит и не изменяет ресурсы в pool. Принимает [Registry] (read-only [res.Entry] walk).

Pipeline:

	fill → materialize → unique.Merge → sdi.Resolve(reg)
	runner / res.GetOneByInterface / res.GetOneByType

Resolve:

	1. collect entries in registration order ([Registry.WalkEntries])
	2. inject — registration order; one-dep match; many → slice in registration order

CheckCycles (optional, not called by Resolve):

	DFS по Deps(); circular → error — for CI/debug when DAG discipline is desired.

Dependency stubs in Deps() []any:

	(*T)(nil)   — exactly one (0 → unresolved, 2+ → ambiguous)
	([]T)(nil)  — many; 0 implementors → empty []T (registration order when non-empty)

Override Replaceable+explicit того же concrete type — [unique.Add] / Merge до Resolve, не sdi.

Org backlog: [github.com/omcrgnt/backlog](https://github.com/omcrgnt/backlog) — [sdi-v21-followups](https://github.com/omcrgnt/backlog/blob/main/items/sdi-v21-followups.md) (DependencyOrder, Many warn, lifecycle lint).
*/
package sdi
