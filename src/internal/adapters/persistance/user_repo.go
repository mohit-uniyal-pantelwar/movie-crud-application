package persistance

type UserRepo struct {
	db *Database
}

func NewUserRepo(d *Database) UserRepo {
	return UserRepo{db: d}
}
