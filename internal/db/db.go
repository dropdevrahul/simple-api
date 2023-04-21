package db

type DBRepo struct {
	User UserRepo
}

func NewDBRepo() *DBRepo {
	return &DBRepo{
		User: &UserPG{
			TableName: "user",
		},
	}
}
