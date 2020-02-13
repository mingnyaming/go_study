package interceptor

import "context"

func Chain(interceptors ...Interceptor) Interceptor {
	n := len(interceptors)

	return func(ctx context.Context, handler Handler) {
		newHandler := func(interceptor Interceptor, handler Handler) Handler {
			return func(ctx context.Context) {
				interceptor(ctx, handler)
			}
		}

		chainedHandler := handler
		for i := n - 1; i >= 0; i-- {
			chainedHandler = newHandler(interceptors[i], chainedHandler)
		}
		chainedHandler(ctx)
	}
}
