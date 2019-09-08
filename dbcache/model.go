package dbcache

type Model interface {
	SqlTemplate() string
	UnmarshalField()
}
