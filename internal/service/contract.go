package service

type DbRepo interface {
	SaveToDB(p1, p2 string) error
}

type ClientC interface {
	SendMessage(message string) (string, error)
}
