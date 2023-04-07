package behavioral

import "fmt"

/*
# Strategy Pattern

## Concept
Turn a set of behaviors into objects and makes them interchangeable inside
original context object.

## Example explanation
The following example consists on the creation of strategy structs that
implements an interface of a Eviction Algorithm. That allows us to change
eviction algorithms on the Context object.
*/

// Strategy interface
type EvictionAlgo interface {
	evict(c *Cache)
}

// Concrete strategy 1
type Fifo struct{}

func (l *Fifo) evict(c *Cache) {
	fmt.Println("Evicting by fifo strategy")
}

// Concrete strategy 2
type Lru struct{}

func (l *Lru) evict(c *Cache) {
	fmt.Println("Evicting by lru strategy")
}

// Concrete strategy 3
type Lfu struct{}

func (l *Lfu) evict(c *Cache) {
	fmt.Println("Evicting by lfu strategy")
}

// Context
type Cache struct {
	storage      map[string]string
	evictionAlgo EvictionAlgo
	capacity     int
	maxCapacity  int
}

func initCache(e EvictionAlgo) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:      storage,
		evictionAlgo: e,
		capacity:     0,
		maxCapacity:  2,
	}
}

func (c *Cache) setEvictionAlgo(e EvictionAlgo) {
	c.evictionAlgo = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) string {
	return c.storage[key]
}

func (c *Cache) evict() {
	c.evictionAlgo.evict(c)
	c.capacity--
}

// Client code
func RunStrategyExample() {
	lfu := &Lfu{}
	cache := initCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")
	cache.add("c", "3")

	lru := &Lru{}
	cache.setEvictionAlgo(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEvictionAlgo(fifo)

	cache.add("e", "5")
	fmt.Println(cache.get("e"))
}
