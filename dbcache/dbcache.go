package dbcache

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

type Item struct {
	Object     Model //真正的数据项  //
	Expiration int64 //生存时间
}

type Cache struct {
	defaultExpiration time.Duration
	items             map[string]Item //缓存数据项存储在map中
	mu                sync.RWMutex    //读写锁
	gcInterval        time.Duration   //过期数据项清理周期
	stopGc            chan bool
	db                *gorm.DB
	typeRegistry      map[string]Model
}

func (c *Cache) RegisterType(typeName string, elem Model) {
	c.typeRegistry[typeName] = elem
}

// 过期缓存数据项清理
func (c *Cache) gcLoop() {
	ticker := time.NewTicker(c.gcInterval)
	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		case <-c.stopGc:
			ticker.Stop()
			return
		}
	}
}

func (c *Cache) del(k string) {
	delete(c.items, k)
}

func (c *Cache) DeleteExpired() {
	now := time.Now().Unix()
	c.mu.Lock()
	defer c.mu.Unlock()

	for k, v := range c.items {
		if v.Expiration > 0 && now > v.Expiration {
			c.del(k)
		}
	}
}

func (c *Cache) set(k string, v Model, d time.Duration) {
	e := time.Now().Add(d).Unix()
	c.mu.Lock()
	c.items[k] = Item{
		Object:     v,
		Expiration: e,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(k string, typeName string) (Model, error) {
	c.mu.RLock()
	item, found := c.items[k]
	if !found {
		c.mu.RUnlock()
		model, err := c.loadDB(k, typeName)
		if err != nil {
			return nil, err
		}
		model.UnmarshalField()
		c.set(k, model, c.defaultExpiration)
		return model, nil
	}
	c.mu.RUnlock()
	return item.Object, nil
}

func (c *Cache) loadDB(k, typeName string) (Model, error) {
	t := c.typeRegistry[typeName]
	tt := reflect.TypeOf(t).Elem()
	model := reflect.New(tt).Interface().(Model)
	err := c.db.Raw(model.SqlTemplate(), k).Scan(model).Error
	return model, err
}
func (c *Cache) Count() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.items)
}

func (c *Cache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items = map[string]Item{}
}

func (c *Cache) Stop() {
	c.stopGc <- true
}

func NewCache(defaultExpiration, gcInterval time.Duration, maxIdleConn, maxOpenConn int, sqlurl string) *Cache {
	c := &Cache{
		defaultExpiration: defaultExpiration,
		gcInterval:        gcInterval,
		items:             map[string]Item{},
		stopGc:            make(chan bool),
		mu:                sync.RWMutex{},
		typeRegistry:      make(map[string]Model),
	}
	db, err := initDB(sqlurl, true, maxIdleConn, maxOpenConn)
	if err != nil {
		panic(err)
	}
	c.db = db
	go c.gcLoop()
	fmt.Println("NewCache ok")
	return c
}
