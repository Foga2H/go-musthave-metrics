package server

type StorageRepo interface {
	Set(key string, value string)
	Get(key string) (string, bool)
}
