package domain

type HashPassword struct {
	EncryptPassword string
}

type CheckHashPassword struct {
	ValidPassword bool
}
