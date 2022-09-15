package data

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOutboundPayload_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    OutboundPayload
		wantErr bool
	}{
		{
			name:    "basic",
			args:    args{[]byte(`{"type": "newHand", "data": {"hand": [{"suit": 1, "value": 3}]}}`)},
			want:    OutboundPayload{"newHand", NewHandEvent{[]Card{{1, 3}}}},
			wantErr: false,
		},
		{
			name:    "invalid type",
			args:    args{[]byte(`{"type": "invalid event", "data": {"hand": [{"suit": 1, "value": 3}]}}`)},
			want:    OutboundPayload{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual OutboundPayload
			err := json.Unmarshal(tt.args.bytes, &actual)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, actual)
			}
		})
	}
}
