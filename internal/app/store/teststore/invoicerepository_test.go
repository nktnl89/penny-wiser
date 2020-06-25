package teststore_test

import (
	"github.com/nktnl89/penny-wiser/internal/app/model"
	"github.com/nktnl89/penny-wiser/internal/app/store"
	"github.com/nktnl89/penny-wiser/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInvoiceRepository_Create(t *testing.T) {
	s := teststore.New()
	i := model.TestInvoice(t)
	assert.NoError(t, s.Invoice().Create(i))
	assert.NotNil(t, i)
}

func TestInvoiceRepository_FindById_should_fail(t *testing.T) {
	s := teststore.New()
	id := -1
	_, err := s.Invoice().FindById(id)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())
}

//func TestInvoiceRepository_FindById(t *testing.T) {
//	db, teardown := sqlstore.TestDB(t, databaseURL)
//	defer teardown("invoices")
//
//	s := sqlstore.New(db)
//
//	_ := s.Invoice().Create(model.TestInvoice(t))
//
//	i, err := s.Invoice().FindById(invoice.ID)
//	assert.NoError(t, err)
//	assert.NotNil(t, i)
//}
