package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isEnv(t *testing.T) {
	type args struct {
		envValue string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test isEnv-true",
			args: args{
				envValue: "${REDIS_ADDRESS_EXPAND}",
			},
			want: true,
		},
		{
			name: "test isEnv-false-1",
			args: args{
				envValue: "REDIS_ADDRESS_EXPAND",
			},
			want: false,
		},
		{
			name: "test isEnv-false-2",
			args: args{
				envValue: "${REDIS_ADDRESS_EXPAND",
			},
			want: false,
		},
		{
			name: "test isEnv-false-3",
			args: args{
				envValue: "REDIS_ADDRESS_EXPAND}",
			},
			want: false,
		},
		{
			name: "test isEnv-false-4",
			args: args{
				envValue: "$REDIS_ADDRESS_EXPAND",
			},
			want: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equalf(t, test.want, isEnv(test.args.envValue), "isEnv(%v)", test.args.envValue)
		})
	}

}
