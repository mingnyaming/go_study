package interceptor

import (
	"context"
	"testing"
)

// todo 動作確認はやってない。
func TestChain(t *testing.T) {
	ctx := context.Background()
	type args struct {
		interceptors []Interceptor
	}
	tests := []struct {
		name string
		args args
		want Interceptor
	}{
		{
			name: "Chain Test",
			args: args{interceptors: []Interceptor{FirstInterceptor(ctx), SecondInterceptor(ctx), ThirdInterceptor(ctx)}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RegisterInterceptor(Chain(tt.args.interceptors...))
			ExcutingHandler(ctx, SampleHandler(ctx))
		})
	}
}
