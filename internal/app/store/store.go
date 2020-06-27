package store

// Store ...
type Store interface {
	Invoice() InvoiceRepository
	Item() ItemRepository
}
