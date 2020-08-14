package simple_lru_cache

type LRU interface {
	Get(key string) (value interface{}, found bool)
	Set(key string, value interface{})
}

type LRUCache struct {
	first *item
	end   *item
	items map[string]*item

	limit int
}

type LRUConfig struct {
	Limit int // item limit
}

type item struct {
	key   string
	value interface{}
	next  *item
	prev  *item
}

func NewLRUCache(config *LRUConfig) *LRUCache {
	return &LRUCache{
		first: nil,
		end:   nil,
		items: make(map[string]*item),

		limit: config.Limit,
	}
}

func (l *LRUCache) Get(key string) (value interface{}, found bool) {
	// check if key exists
	if item, ok := l.items[key]; ok {
		l.promoteItem(item)
		return item.value, true
	}
	return nil, false
}

func (l *LRUCache) Set(key string, value interface{}) {
	// create item
	item := &item{
		key:   key,
		value: value,
		next:  nil,
		prev:  nil,
	}

	// check if first is empty else add to the end
	if l.first == nil {
		l.first = item
	} else if l.end == nil {
		l.end = item
		l.end.prev = l.first
		l.first.next = l.end
	} else {
		if len(l.items) == l.limit {
			l.evict()
		}
		item.prev = l.end
		item.prev.next = item
		l.end = item
	}

	// set item to map
	l.items[key] = item
}

func (l *LRUCache) evict() {
	delete(l.items, l.first.key)
	l.first = l.first.next
	l.first.prev = nil
}

// promoteItem moves item to end and replaces references to neighbors
func (l *LRUCache) promoteItem(i *item) {
	if i != l.end {
		if i != l.first {
			i.prev.next = i.next
		} else {
			l.first = i.next
		}
		i.next.prev = i.prev
		l.end.next = i
		i.prev = l.end
		l.end = i
	}
}
