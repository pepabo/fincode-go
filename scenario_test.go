package fincode

import (
	"context"
	"testing"

	"github.com/k1LoW/runn"
)

func TestScenario(t *testing.T) {
	ctx := context.Background()
	ts := NewPaymentServer(t)
	opts := []runn.Option{
		runn.T(t),
		runn.SkipIncluded(true),
		runn.Var("endpoint", ts.URL),
	}
	o, err := runn.Load("testdata/scenarios/**/*.yml", opts...)
	if err != nil {
		t.Fatal(err)
	}

	if err := o.RunN(ctx); err != nil {
		t.Error(err)
	}

	t.Logf("%d scenarios", o.Result().Total.Load())
}
