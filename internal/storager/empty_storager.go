package storager

// заглушка когда без хранилища
type EmptyStorager struct {
}

// загрузить из базы
func (one *EmptyStorager) Load(shorten *[]*Shorten) (*[]*Shorten, error) {
	return shorten, nil
}

// положить в базу
func (one *EmptyStorager) StorageWrite(short, origin string, UUID int) error {
	return nil
}

// удалить из базы
func (one *EmptyStorager) Delete(short interface{}) error {
	return nil
}

// конструктор
func NewEmptyStorager() (*EmptyStorager, error) {
	store := &EmptyStorager{}
	return store, nil

}
