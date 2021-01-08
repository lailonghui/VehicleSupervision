package dataloader

import (
	"context"
	"github.com/gin-gonic/gin"
	"reflect"
	"sync"
)

const DATA_LOADER_CONTEXT_KEY = "C_DATALOADER"

// dataloader集合
type Loaders struct {
	sync.Mutex
	loaderMap map[string]interface{}
}

func NewLoaders() *Loaders {
	return &Loaders{
		Mutex:     sync.Mutex{},
		loaderMap: make(map[string]interface{}, 0),
	}
}

//GetLoader 获取dataloader
func (t *Loaders) GetLoader(dest interface{}) interface{} {
	iType := reflect.TypeOf(dest).Elem()
	typeStr := iType.String()
	i, ok := t.loaderMap[typeStr]
	if !ok {
		t.Mutex.Lock()
		i, ok = t.loaderMap[typeStr]
		if !ok {
			v := reflect.New(iType)
			d := v.MethodByName("NewLoader").Call([]reflect.Value{})[0].Interface()
			t.loaderMap[typeStr] = d
			i = d
		}
		t.Mutex.Unlock()
	}
	return i
}

// dataloader 中间件
func DataloaderMiddle(contextKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), contextKey, NewLoaders())
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}

}
