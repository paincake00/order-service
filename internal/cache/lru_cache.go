package cache

import (
	"container/list"

	"github.com/paincake00/order-service/internal/domain/model"
)

type LRUCache struct {
	Capacity   int
	cache      map[string]*list.Element
	linkedList *list.List
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity:   capacity,
		cache:      make(map[string]*list.Element),
		linkedList: list.New(),
	}
}

func (c *LRUCache) Get(uid string) (*model.OrderModel, bool) {
	if el, ok := c.cache[uid]; ok {
		c.linkedList.MoveToFront(el)
		return el.Value.(*model.OrderModel), true
	}
	return &model.OrderModel{}, false
}

func (c *LRUCache) Put(order *model.OrderModel) {
	if el, ok := c.cache[order.OrderUID]; ok {
		c.linkedList.MoveToFront(el)
		return
	}
	newEl := c.linkedList.PushFront(order)
	c.cache[order.OrderUID] = newEl

	if c.linkedList.Len() > c.Capacity {
		old := c.linkedList.Back()
		if old != nil {
			c.linkedList.Remove(old)
			delete(c.cache, old.Value.(*model.OrderModel).OrderUID)
		}
	}
}
