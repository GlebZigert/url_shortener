package storager

type EmptyStorager struct {
}

func (one *EmptyStorager) Load(shorten *[]*Shorten) error {
	return nil
}

func (one *EmptyStorager) StorageWrite(short, origin string, UUID int) error {
	return nil
}

func (one *EmptyStorager) Delete(listID []int) {

}

func (one *EmptyStorager) Init() error {

	return nil

}
