package entity

type User struct {
	UserID         string
	Email          string
	EmailToUpdate  string
	HashedPassword string // usecase内でHashedPasswordを生成し、設定する
	IsDeleted      bool
}