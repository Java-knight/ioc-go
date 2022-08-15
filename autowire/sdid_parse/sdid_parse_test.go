package sdid_parse

import (
	"github.com/stretchr/testify/assert"
	"ioc-go/autowire"
	"testing"
)

func TestGetDefaultSDIDParser(t *testing.T) {
	t.Run("Get Default SDID Parser equal", func(t *testing.T) {
		got1 := GetDefaultSDIDParser()
		got2 := GetDefaultSDIDParser()
		assert.Equal(t, got1, got2)
	})

	t.Run("Get Default SDID Parser not nil", func(t *testing.T) {
		assert.NotNil(t, GetDefaultSDIDParser())
	})
}

func Test_defaultSDIDParser_Parse(t *testing.T) {
	type args struct {
		fieldInfo *autowire.FieldInfo
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test default sdid parse normal interface field info",
			args: args{
				fieldInfo: &autowire.FieldInfo{
					FieldName: "MyRedis",
					FieldType: "Redis",
					TagKey:    "normal",
					TagValue:  "Impl",
				},
			},
			want:    "Impl",
			wantErr: false,
		},
		{
			name: "test default sdid parse normal struct field info",
			args: args{
				fieldInfo: &autowire.FieldInfo{
					FieldName: "MyRedis",
					FieldType: "",
					TagKey:    "normal",
					TagValue:  "StructImpl",
				},
			},
			want:    "StructImpl",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &defaultSDIDParse{}
			got, err := p.Parse(tt.args.fieldInfo)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
