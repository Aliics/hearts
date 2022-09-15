package data

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInboundPayload_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    InboundPayload
		wantErr bool
	}{
		{
			name:    "basic",
			args:    args{[]byte(`{"type": "playCard", "data": {"card": {"suit": 1, "value": 3}}}`)},
			want:    InboundPayload{"playCard", PlayCardEvent{Card: Card{1, 3}}},
			wantErr: false,
		},
		{
			name:    "invalid type",
			args:    args{[]byte(`{"type": "invalid event", "data": {"card": {"suit": 1, "value": 3}}}`)},
			want:    InboundPayload{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var actual InboundPayload
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
