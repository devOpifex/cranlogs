package data_test

import (
	"fmt"
	"testing"

	"github.com/devOpifex/cranlogs/internal/data"
	"github.com/stretchr/testify/assert"
)

func TestPeriod(t *testing.T) {
	type test struct {
		input   string
		want    data.Period
		wantErr error
	}

	tests := []test{
		{input: "last-week", want: data.Period("last-week"), wantErr: nil},
		{input: "2022-01-01", want: data.Period("2022-01-01"), wantErr: nil},
		{input: "2022-01-01:2022-01-05", want: data.Period("2022-01-01:2022-01-05"), wantErr: nil},
		{input: "last-day", want: data.Period(""), wantErr: fmt.Errorf("invalid period: last-day")},
	}
	for _, testVal := range tests {
		p, err := data.NewPeriod(testVal.input)
		assert.Equal(t, testVal.want, p)
		assert.Equal(t, testVal.wantErr, err)
	}
}
