package ports

type Password interface {
	HashPassword(pw string) (string, error)
	VerifyPassword(pw, hashed string) (bool, error)
}
