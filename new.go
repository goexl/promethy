package promethy

import (
	"sync"

	"github.com/goexl/promethy/internal/builder"
)

var once sync.Once

// New 创建普罗米修斯
// 使用单例模式，多次创建只会创建一个实例
func New() (promethy *builder.Promethy) {
	once.Do(func() {
		promethy = builder.NewPromethy()
	})

	return
}
