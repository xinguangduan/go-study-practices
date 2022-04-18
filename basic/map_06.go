package main

import (
	"fmt"
	"sync"
)

// 使用场景
//sync.Map更适合读多更新多而插入新值少的场景（appendOnly模式，尤其是key存一次，多次读而且不删除的情况），
//因为在key存在的情况下读写删操作可以不用加锁直接访问readOnly不适合反复插入与读取新值的场景，
//因为这种场景会频繁操作dirty，需要频繁加锁和更新read【此场景github开源库orcaman/concurrent-map更合适】
func main() {
	m := sync.Map{}
	m.Store("key", 1)
	fmt.Println(m.Load("key"))
}
