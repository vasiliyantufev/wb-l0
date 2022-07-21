package cache

import (
	"sync"
)

type Cache struct {
	mx   sync.RWMutex
	Data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		Data: make(map[string]string),
	}
}

func (c *Cache) PutOrder(id string, o string) {
	c.mx.Lock()
	defer c.mx.Unlock()
	c.Data[id] = o
}

func (c *Cache) GetOrder(id string) (o string, b bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()
	o, b = c.Data[id]
	return
}
