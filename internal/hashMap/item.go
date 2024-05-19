package hashMap

type Item struct {
	key  string
	val  interface{}
	next *Item
}

func NewItem(key string, val interface{}) *Item {
	return &Item{
		key:  key,
		val:  val,
		next: nil,
	}
}
