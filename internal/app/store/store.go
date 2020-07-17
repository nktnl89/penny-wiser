package store

// Store ...
type Store interface {
	Invoice() InvoiceRepository
	Item() ItemRepository
	Plan() PlanRepository
	PlanItem() PlanItemRepository
	Entry() EntryRepository
}
