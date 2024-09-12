package repositories

type OddRepository interface {
	GetOdd() (int, error)
}

type oddRepository struct {
	db string
}

func NewOddRepository(db string) OddRepository {
	return &oddRepository{db: db}
}

func (r *oddRepository) GetOdd() (int, error) {
	return 10, nil
}
