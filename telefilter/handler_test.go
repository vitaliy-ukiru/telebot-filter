package telefilter

import (
	"testing"

	tb "gopkg.in/telebot.v3"
)

func TestRawHandler_Check(t *testing.T) {

	trueFilter := func(tb.Context) bool { return true }
	falseFilter := func(tb.Context) bool { return false }

	tests := []struct {
		name    string
		filters []Filter
		want    bool
	}{
		{
			name: "without filters",
			want: true,
		},
		{
			name:    "all positive",
			filters: []Filter{trueFilter, trueFilter},
			want:    true,
		},
		{
			name:    "negative handler",
			filters: []Filter{trueFilter, trueFilter, falseFilter, trueFilter},
			want:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := RawHandler{
				Filters: tt.filters,
			}
			if got := h.Check(nil); got != tt.want {
				t.Errorf("Check() = %v, want %v", got, tt.want)
			}
		})
	}
}
