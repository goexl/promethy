package promethy

import (
	"sync"
)

var (
	once sync.Once
	_    = New
)

// New 创建普罗米修斯
// 使用单例模式，多次创建只会创建一个实例
func New() (builder *builder) {
	once.Do(func() {
		builder = newBuilder()
	})

	return
}
