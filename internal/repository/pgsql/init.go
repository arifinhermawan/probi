package pgsql

type pgsqlProvider interface {
}

type DBRepo struct {
	db pgsqlProvider
}

func NewDBRepository(db pgsqlProvider) *DBRepo {
	return &DBRepo{
		db: db,
	}
}
