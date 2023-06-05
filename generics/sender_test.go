package generics

import (
	"testing"
)

func TestSender_SendInt(t *testing.T) {
	type args struct {
		msg int
	}
	tests := []struct {
		name string
		s    *Sender
		args args
		want int
	}{
		{
			name: "1",
			s:    &Sender{},
			args: args{
				msg: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SendInt(tt.args.msg); got != tt.want {
				t.Errorf("Sender.SendInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSender_SendString(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		s    *Sender
		args args
		want string
	}{
		{
			name: "hello",
			s:    &Sender{},
			args: args{
				msg: "hello",
			},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.SendString(tt.args.msg); got != tt.want {
				t.Errorf("Sender.SendString() = %v, want %v", got, tt.want)
			}
		})
	}
}
