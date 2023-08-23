package string

import "testing"

func Test_toDigit(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		wantRes int64
	}{
		{"test", args{"三千五百四十二万五千三百六十八亿七千九百六十万零五千二百三十八"}, 3542536879605238},
		{"test", args{"三十亿零五千万零三"}, 3050000003},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := toDigit(tt.args.str); gotRes != tt.wantRes {
				t.Errorf("toDigit() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
