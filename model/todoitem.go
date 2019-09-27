package model

type TODOItem struct {
	ID             string
	OwnerID        string
	TopicID        string
	FatherID       string
	Content        string
	Status         uint32
	CompletionTime uint64
	DeletedTime    uint64
}
