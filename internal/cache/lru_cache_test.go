package cache

import (
	"testing"

	"github.com/paincake00/order-service/internal/domain/model"
)

func TestLRUCache_PutAndGet(t *testing.T) {
	c := NewLRUCache(3)

	o := &model.OrderModel{
		OrderUID: "1",
	}

	c.Put(o)

	res, ok := c.Get("1")

	if !ok {
		t.Errorf("should get order with uid=1")
	}

	if res.OrderUID != o.OrderUID {
		t.Errorf("expected UID %s, got %s", o.OrderUID, res.OrderUID)
	}
}

func TestLRUCache_RemoveOldest(t *testing.T) {
	c := NewLRUCache(2)

	o1 := &model.OrderModel{OrderUID: "1"}
	o2 := &model.OrderModel{OrderUID: "2"}
	o3 := &model.OrderModel{OrderUID: "3"}

	c.Put(o1)
	c.Put(o2)

	_, _ = c.Get("1") // первый теперь uid=1

	c.Put(o3)

	if _, ok := c.Get("2"); ok {
		t.Errorf("order with uid=2 should not exist")
	}

	if _, ok := c.Get("1"); !ok {
		t.Errorf("order with uid=1 should exist")
	}

	if _, ok := c.Get("3"); !ok {
		t.Errorf("order with uid=3 should exist")
	}
}
