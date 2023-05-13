package mtClient

import "testing"

func TestGetType(t *testing.T) {
	type args struct {
		destination string
	}
	tests := []struct {
		name string
		args args
		want DestinationType
	}{
		{
			name: "email",
			args: args{
				destination: "my@email.com",
			},
			want: DestinationUserEmail,
		},
		{
			name: "channel",
			args: args{
				destination: "~channelName",
			},
			want: DestinationChannel,
		},
		{
			name: "userName",
			args: args{
				destination: "userName",
			},
			want: DestinationUserName,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetType(tt.args.destination); got != tt.want {
				t.Errorf("GetType() = %v, want %v", got, tt.want)
			}
		})
	}
}
