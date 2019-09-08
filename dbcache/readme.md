# 基于mysql到内存的缓存库
## 背景
- 很多业务上需要读取db里面的数据详情，这些数据已经存在db，会修改但不会删除,程序为了加速访问需要缓存到内存，业务上也允许内存和db 不一致,设计一个通用的dbcache 由此而来


##  最佳实践
```
	cache := NewCache(24*time.Hour, 10*time.Minute, 10, 10, "root:test@tcp(127.0.0.1:3306)/test?parseTime=True&loc=Local&multiStatements=true&charset=utf8")
	c := Cri{}
	cache.RegisterType("cri", &c)
	data, err := cache.Get("1", "cri")
	fmt.Println(data, err, "=======================")
```

- 1. 首先new 一个cache

- 2. 需要实现 Model 接口
>	SqlTemplate() string  返回查询sql 语句模板 
>	UnmarshalField() Model 对象里面有些字段可能不是db直接查出而是序列号而来 
- 3. 注册model 
> func (c *Cache) RegisterType(typeName string, m Model) typeName 为类型名称,m 为实现的model(传指针)

- 4. 获取值
> func (c *Cache) Get(k string,typeName) (Model, error) typeName 为类型名称 k 为参数

## 注意
1. 此cache 返回data为只读，不可修改
2. mysql 需存在此model 否则会每次穿透db
