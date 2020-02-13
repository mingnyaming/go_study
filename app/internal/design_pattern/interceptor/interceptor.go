package interceptor

import "context"

type Handler func(ctx context.Context)

type Interceptor func(ctx context.Context, handler Handler)

var interceptor Interceptor

func ExcutingHandler(ctx context.Context, handler Handler) {
	if interceptor == nil {
		handler(ctx)
	} else {
		interceptor(ctx, handler)
	}
}

func FirstInterceptor(ctx context.Context) Interceptor {
	return func(ctx context.Context, handler Handler) {
		println("FirstInterceptor: before Handler")
		handler(ctx)
		println("FirstInterceptor: after Handler")
	}
}

func SecondInterceptor(ctx context.Context) Interceptor {
	return func(ctx context.Context, handler Handler) {

		println("SecondInterceptor: before Handler")
		handler(ctx)
		println("SecondInterceptor: after Handler")
	}
}

func ThirdInterceptor(ctx context.Context) Interceptor {
	return func(ctx context.Context, handler Handler) {

		println("ThirdInterceptor: before Handler")
		handler(ctx)
		println("ThirdInterceptor: after Handler")
	}
}

func SampleHandler(ctx context.Context) Handler {
	return func(ctx context.Context) {

		println("@: This is main function")
	}
}

func RegisterInterceptor(inputInterceptor Interceptor) {
	interceptor = inputInterceptor
}

func SampleInterceptor(ctx context.Context) Interceptor {
	return func(ctx context.Context, handler Handler) {
		println("1: before Handler")
		handler(ctx)
		println("1: after Handler")
	}
}
