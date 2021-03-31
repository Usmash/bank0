package types

// Money for dirams, sents, kopeyks
type Money int64

// Category describes payment category
type Category string

// Payment describes structure for payments
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}