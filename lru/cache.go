package lru

func New[K comparable, V any](capacity int) Cache[K, V] {
	if capacity <= 0 {
		panic("lru: capacity must be greater than 0")
	}

	return &cache[K, V]{
		capacity:      capacity,
		lookup:        make(map[K]*node[V], capacity),
		reverseLookup: make(map[*node[V]]K, capacity),
	}
}

type Cache[K comparable, V any] interface {
	Update(K, V)
	Get(K) (V, bool)
}

type cache[K comparable, V any] struct {
	length        int
	capacity      int
	head          *node[V]
	tail          *node[V]
	lookup        map[K]*node[V]
	reverseLookup map[*node[V]]K
}

type node[T any] struct {
	val  T
	next *node[T]
	prev *node[T]
}

func (c *cache[K, V]) Update(key K, val V) {
	n, ok := c.lookup[key]

	if ok {
		c.detach(n)
		c.prepend(n)
		n.val = val
		return
	}

	n = &node[V]{val: val}
	c.length++
	c.prepend(n)
	c.trimCache()
	c.lookup[key] = n
	c.reverseLookup[n] = key
}

func (c *cache[K, V]) Get(key K) (val V, ok bool) {
	n, ok := c.lookup[key]

	if !ok {
		return val, false
	}

	c.detach(n)
	c.prepend(n)

	return n.val, true
}

func (c *cache[K, V]) detach(n *node[V]) {
	if n.prev != nil {
		n.prev.next = n.next
	}

	if n.next != nil {
		n.next.prev = n.prev
	}

	if c.head == n {
		c.head = n.next
	}

	if c.tail == n {
		c.tail = n.prev
	}

	n.prev = nil
	n.next = nil
}

func (c *cache[K, V]) prepend(n *node[V]) {
	if c.head == nil {
		c.head = n
		c.tail = n
		return
	}

	n.next = c.head
	c.head.prev = n
	c.head = n
}

func (c *cache[K, V]) trimCache() {
	if c.length <= c.capacity {
		return
	}

	tail := c.tail
	c.detach(tail)
	key := c.reverseLookup[tail]
	delete(c.lookup, key)
	delete(c.reverseLookup, tail)
	c.length--
}
