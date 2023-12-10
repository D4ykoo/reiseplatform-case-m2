package adapter

type CrudRepository[T any] interface {
	create(T)
	update(T)
	get(uuid string) T
	getBy(statemant string, parameter ...string) []T
	delete(T) bool
}
