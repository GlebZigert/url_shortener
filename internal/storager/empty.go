package storager

type EmptyStorager struct {
}

func (one *EmptyStorager) Load(mapa *map[string]string) error {
	return nil
}

func (one *EmptyStorager) StorageWrite(short, origin string, id, UUID int) error {
	return nil
}

func (one *EmptyStorager) Delete(str string) {

}

func (one *EmptyStorager) Init() error {

	return nil

}
