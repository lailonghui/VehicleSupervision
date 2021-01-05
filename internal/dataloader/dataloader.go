package dataloader

// dataloader接口
type Dataloader interface {
	NewLoader() interface{}
}
