package repository

type Repository struct {
	storage map[string]string
}

func NewRepository() *Repository {
	s := make(map[string]string, 0)
	return &Repository{storage: s}
}

func (r *Repository) SaveToDB(p1, p2 string) error {
	r.storage[p1] = p2
	return nil
}
