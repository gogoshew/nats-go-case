package repository

type Repository struct {
	Db    DataBase
	Cache Cacher
}

func NewRepository(db DataBase, cache Cacher) *Repository {
	return &Repository{Db: db, Cache: cache}
}
