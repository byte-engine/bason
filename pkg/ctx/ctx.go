package ctx

import (
	"github.com/byte-engine/bason/pkg/lang"
)

type Ctx interface {
	Put(key any, val any) Ctx
	Value(key any) any
	Remove(key any) Ctx
	Contain(key any) bool
}

// region ctx

// ctx Ctx基础实现
type ctx struct {
	Ctx
	parent Ctx
	data   *lang.ConcurrentMap[any, any]
	err    error
}

func (c *ctx) Value(key any) any {
	var (
		val    any
		ok     bool
		parent Ctx
	)
	if val, ok = c.data.Get(key); ok {
		return val
	}
	parent = c.getParent()
	if parent != nil {
		return c.parent.Value(key)
	}
	return nil
}

func (c *ctx) Push(key any, value any) Ctx {
	c.data.Set(key, value)
	return c
}

func (c *ctx) getParent() Ctx {
	return c.parent
}

func (c *ctx) Contain(key any) bool {
	return c.data.Has(key)
}

func (c *ctx) Remove(key any) Ctx {
	c.data.Delete(key)
	return c
}

// endregion

// region immutableCtx

// immutableCtx
type immutableCtx struct {
	ctx
}

func (c *immutableCtx) Push(key any, value any) Ctx {
	return c
}

func (c *immutableCtx) Delete(key any) Ctx {
	return c
}

//endregion

// region global functions

// region Root functions

// Root no parent Ctx
func Root() Ctx {
	return &ctx{
		data: lang.NewConcurrentMap[any, any](128),
	}
}

func RootFrom(data map[any]any) Ctx {
	var (
		target = &ctx{
			data: lang.NewConcurrentMap[any, any](128),
		}
		tmpData *lang.ConcurrentMap[any, any]
	)
	tmpData = target.data
	if data != nil && len(data) > 0 {
		for key, val := range data {
			tmpData.Set(key, val)
		}
	}
	return target
}

func RootImmutable(data map[any]any) Ctx {
	var (
		target = &immutableCtx{
			ctx: ctx{
				data: lang.NewConcurrentMap[any, any](128),
			},
		}
		tmpData *lang.ConcurrentMap[any, any]
	)
	tmpData = target.data
	if data != nil && len(data) > 0 {
		for key, val := range data {
			tmpData.Set(key, val)
		}
	}
	return target
}

// endregion

// region New functions

// NewWithParent new Ctx from parent
func NewWithParent(parent Ctx) Ctx {
	return &ctx{
		parent: parent,
		data:   lang.NewConcurrentMap[any, any](128),
	}
}

func NewFromWithParent(data map[any]any, parent Ctx) Ctx {
	var (
		target = &ctx{
			parent: parent,
			data:   lang.NewConcurrentMap[any, any](128),
		}
		tmpData *lang.ConcurrentMap[any, any]
	)
	tmpData = target.data
	if data != nil && len(data) > 0 {
		for key, val := range data {
			tmpData.Set(key, val)
		}
	}
	return target
}

func NewImmutableFromParent(data map[any]any, parent Ctx) Ctx {
	var (
		target = &immutableCtx{
			ctx: ctx{
				parent: parent,
				data:   lang.NewConcurrentMap[any, any](128),
			},
		}
		tmpData *lang.ConcurrentMap[any, any]
	)
	tmpData = target.data
	if data != nil && len(data) > 0 {
		for key, val := range data {
			tmpData.Set(key, val)
		}
	}
	return target
}

func NewMappingFrom(parent Ctx, keyMapping map[any]any) Ctx {
	var (
		target = &ctx{
			parent: parent,
			data:   lang.NewConcurrentMap[any, any](128),
		}
		tmpData *lang.ConcurrentMap[any, any]
	)
	tmpData = target.data
	for newKey, fromKey := range keyMapping {
		tmpData.Set(newKey, parent.Value(fromKey))
	}
	return target
}

func NewImmutableMappingFrom(parent Ctx, keyMapping map[any]any) Ctx {
	var (
		target = &immutableCtx{
			ctx: ctx{
				parent: parent,
				data:   lang.NewConcurrentMap[any, any](128),
			},
		}
		tmpData *lang.ConcurrentMap[any, any]
	)
	tmpData = target.data
	for newKey, fromKey := range keyMapping {
		tmpData.Set(newKey, parent.Value(fromKey))
	}
	return target
}

// endregion

// endregion
