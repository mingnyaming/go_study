package interceptor

import (
	"context"
	"testing"
)

// todo 動作確認はやってない。
func TestInterceptor(t *testing.T) {
	ctx := context.Background()
	type args struct {
		interceptor Interceptor
		handler     Handler
	}
	tests := []struct {
		name string
		args args
		want Interceptor
	}{
		{
			name: "interceptor Test: Sample",
			args: args{
				interceptor: SampleInterceptor(ctx),
				handler:     SampleHandler(ctx),
			},
		},
	}
	for _, tt := range tests {
		ic := tt.args.interceptor
		hd := tt.args.handler
		ic(ctx, hd)
	}
}
