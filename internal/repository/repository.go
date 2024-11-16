package repository

type Repository struct {
	dataMap map[struct{}]int
}

func New() *Repository {
	return &Repository{
		dataMap: make(map[struct{}]int),
	}
}

func (r *Repository) Save(chislo int) {
	r.dataMap[struct{}{}] = chislo
}
