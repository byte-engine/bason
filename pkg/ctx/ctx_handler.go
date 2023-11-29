package ctx

type CtxConsumer func(c Ctx)

func AdaptSupplier[R any](supplier func(c Ctx) R, resultKey string) CtxConsumer {
	return func(c Ctx) {
		result := supplier(c)
		c.Put(resultKey, result)
	}
}

func AdaptFunction[A any, R any](aSupplier func(c Ctx) A, fn func(arg A) R, resultKey string) CtxConsumer {
	return func(c Ctx) {
		arg := aSupplier(c)
		result := fn(arg)
		c.Put(resultKey, result)
	}
}
