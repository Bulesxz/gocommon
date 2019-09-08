package dbcache

import (
	"fmt"
	"testing"
	"time"
)

type Cri struct {
	ID          int `gorm:"-"`
	RawData     string
	Data1       string `gorm:"-"`
	SQLTemplate string `gorm:"-"`
}

func (m *Cri) SqlTemplate() string {
	return "select raw_data from cir where cri.id=?"
}

func (m *Cri) UnmarshalField() {
}

//go test -v
func TestA(t *testing.T) {
	cache := NewCache(24*time.Hour, 10*time.Minute, 10, 10, "root:test@tcp(127.0.0.1:3306)/test?parseTime=True&loc=Local&multiStatements=true&charset=utf8")
	c := Cri{}
	cache.RegisterType("cri", &c)
	data, err := cache.Get("1", "cri")
	fmt.Println(data, err, "=======================")
}

//go test -v -run=TestB
func TestB(t *testing.T) {
	cache := NewCache(5*time.Second, 1*time.Second, 10, 10, "root:test@tcp(127.0.0.1:3306)/test?parseTime=True&loc=Local&multiStatements=true&charset=utf8")
	c := Cri{}
	cache.RegisterType("cri", &c)
	data, err := cache.Get("1", "cri")
	fmt.Println(data, err, "================find=======")
	time.Sleep(2 * time.Second)
	data, err = cache.Get("1", "cri")
	fmt.Println(data, err, "==========find=============")

	time.Sleep(5 * time.Second)
	data, err = cache.Get("1", "cri")
	fmt.Println(data, err, "==========no find=============")
}

//go test -bench=BenchmarkGet -benchtime=20s
func BenchmarkGet(b *testing.B) {
	cache := NewCache(24*time.Hour, 10*time.Minute, 10, 10, "root:test@tcp(127.0.0.1:3306)/test?parseTime=True&loc=Local&multiStatements=true&charset=utf8")
	c := Cri{}
	cache.RegisterType("cri", &c)
	for i := 0; i < b.N; i++ { //use b.N for looping
		cache.Get("1", "cri")
	}
	fmt.Println("Count:", cache.Count())
}
