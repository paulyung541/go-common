package collection

import (
	"fmt"
	"strings"
)

// 非线程安全Set，可用于去重
type Set struct {
	items map[interface{}]struct{}
}

// 新建 Set 集合
func NewSet(items ...interface{}) *Set {
	set := &Set{
		items: make(map[interface{}]struct{}, 10),
	}

	for _, item := range items {
		set.items[item] = struct{}{}
	}

	return set
}

// Add will add the provided items to the set.
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = struct{}{}
	}
}

func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

// 将 Set 转为切片List
func (set *Set) List() []interface{} {
	l := make([]interface{}, 0, len(set.items))
	for item := range set.items {
		l = append(l, item)
	}
	return l
}

// 返回 Set 元素个数
func (set *Set) Len() int {
	return len(set.items)
}

// 清空 Set
func (set *Set) Clear() {
	set.items = map[interface{}]struct{}{}
}

// items 是否全部都存在于 Set 集合中
func (set *Set) All(items ...interface{}) bool {
	for _, item := range items {
		if _, ok := set.items[item]; !ok {
			return false
		}
	}
	return true
}

// items 是否任意一个存在于 Set 集合中
func (set *Set) Any(items ...interface{}) bool {
	for _, item := range items {
		if _, ok := set.items[item]; ok {
			return true
		}
	}
	return false
}

// 打印专用
func (set *Set) String() string {
	builder := new(strings.Builder)
	builder.WriteString(`Set[`)

	strs := make([]string, 0, set.Len())
	for item := range set.items {
		strs = append(strs, fmt.Sprintf("\"%+v\"", item))
	}

	builder.WriteString(strings.Join(strs, ","))
	builder.WriteString(`]`)
	return builder.String()
}
