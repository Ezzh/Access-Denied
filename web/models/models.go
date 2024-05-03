package models

type SCP struct {
	Name            string
	DescryptionPath string
	ImagePath       string
	Author          string
	IsSecret        bool
}

type User struct {
	Username string
	Password string
}
