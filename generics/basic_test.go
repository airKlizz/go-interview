package generics

import "testing"

func Test_maxInt(t *testing.T) {
	type args struct {
		c1 int
		c2 int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1 2",
			args: args{
				c1: 1,
				c2: 2,
			},
			want: 2,
		},
		{
			name: "2 1",
			args: args{
				c1: 2,
				c2: 1,
			},
			want: 2,
		},
		{
			name: "1 1",
			args: args{
				c1: 1,
				c2: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxInt(tt.args.c1, tt.args.c2); got != tt.want {
				t.Errorf("maxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxFloat32(t *testing.T) {
	type args struct {
		c1 float32
		c2 float32
	}
	tests := []struct {
		name string
		args args
		want float32
	}{
		{
			name: "1 2",
			args: args{
				c1: 1.,
				c2: 2,
			},
			want: 2,
		},
		{
			name: "2 1",
			args: args{
				c1: 2,
				c2: 1,
			},
			want: 2,
		},
		{
			name: "1 1",
			args: args{
				c1: 1,
				c2: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxFloat32(tt.args.c1, tt.args.c2); got != tt.want {
				t.Errorf("maxInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
