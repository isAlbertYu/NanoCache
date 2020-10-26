package lru

import "container/list"

type Cache struct {
	maxBytes         int64                         //允许使用的最大内存
	nBytes           int64                         //当前已使用的内存
	doubleLinkedList *list.List                    //go标准库中的双向链表
	cacheMap         map[string]*list.Element      //map中存储key与对应的链表节点，节点值为entry类型
	OnEvicted        func(key string, value Value) //当某条记录被淘汰是的回调函数
}

//双向链表中的节点值为entry类型
type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int //获取值所占的内存大小
}

//根据key，获取value
//当某个value被访问到时，将其在链表中的位置移动至队头
func (c *Cache) Get(key string) (Value, bool) {
	if element, ok := c.cacheMap[key]; ok {
		c.doubleLinkedList.MoveToFront(element) //移动至队头
		kv := element.Value.(*entry)
		return kv.value, true
	}
	return nil, false
}
