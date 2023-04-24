package db

type DBRepo struct {
	User UserRepo
  UserToken UserTokenRepo
}

func NewDBRepo() *DBRepo {
	return &DBRepo{
		User: &UserPG{
			TableName: "users",
		},
    UserToken: &UserTokenPG{
        TableName: "user_tokens",
    },
	}
}
