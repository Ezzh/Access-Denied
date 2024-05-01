package models

type SCP struct {
	Name            string
	DescryptionPath string
	ImagePath       string
	IsSecret        bool
}

type User struct {
	Username string
	Password string
}
