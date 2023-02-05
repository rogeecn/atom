package dictionary

type DictInfo struct {
	ID          uint64
	Name        string
	Alias       string
	Description string
	Items       []*DictItem
}

type DictItem struct {
	Label string
	Value string
}
