package entity

type Cart struct {
	ID    uint
	Items []Item
}

type Item struct {
	ProductID uint
	Quantity  int
}
