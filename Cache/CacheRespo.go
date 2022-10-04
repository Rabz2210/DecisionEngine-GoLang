package Cache

type AbsCache interface {
	Add(string)
	Get(string) bool
	Delete(string)
}
