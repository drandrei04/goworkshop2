package persistence

import (
	"errors"
	"goworkshop/model"
)

func (store *GormDataStore) CreateBook(book *model.Book) error {
	return store.DBInstance.Create(&book).Error
}

func (store *GormDataStore) GetBook(uuid string) (model.Book, error) {
	var book model.Book
	err := store.DBInstance.Preload("Author").Where(&model.Book{Entity: model.Entity{UUID: uuid}}).First(&book).Error
	return book, err
}

func (store *GormDataStore) GetBooks() ([]model.Book, error) {
	var books []model.Book
	err := store.DBInstance.
	//fetch the author
		Preload("Author").
		Find(&books).Error
	return books, err
}

func (store *GormDataStore) DeleteBook(uuid string) error {
	if isEmptyUUID(uuid) {
		return errors.New("Please supply a primary key when attempting to delete an object.")
	}
	db := store.DBInstance.
		Delete(model.Book{}, whereBookUUIDEquals(uuid))

	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return errors.New("No rows have been deleted for id = " + uuid)
	} else if db.RowsAffected > 1 {
		return errors.New("Multiple rows have been deleted for uuid = " + uuid)
	}
	return nil
}

func (store *GormDataStore) UpdateBook(book *model.Book) error {
	return store.DBInstance.Save(&book).Error
}

//creates the GORM where clause used to identify a book by uuid
func whereBookUUIDEquals(uuid string) (interface{}) {
	return model.Book{Entity: model.Entity{UUID: uuid}}
}
