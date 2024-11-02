package service

type ItemNotFound struct {
	Item string
}

func (e *ItemNotFound) Error() string {
	return e.Item + " not found"
}

type ItemAlreadyExists struct {
	Item string
}

func (e *ItemAlreadyExists) Error() string {
	return e.Item + " already exists"
}

type InvalidDatabaseItem struct{}

func (e *InvalidDatabaseItem) Error() string {
	return "Invalid database item"
}
