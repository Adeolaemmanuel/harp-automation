package controls

import (
	"testing"
)

var control = Controls{}

func TestAndCombinator(t *testing.T) {

	data := Controls{
		Id:         "0a6ccc76-0786-4bd5-99c2-3d516bca165a",
		Combinator: "and",
		Rules: []Rules{
			{
				Id:       "2e4fb517-4db7-46fc-90a7-542c8ec08f63",
				Field:    "Steve",
				Operator: "=",
				Value:    "Steve",
			},
			{
				Id:       "13a7bab3-c785-4403-8b2e-39c27cccc703",
				Field:    40.10,
				Operator: ">",
				Value:    20.20,
			},
			{
				Id:       "13a7bab3-c785-4403-8b2e-39c27cccc703",
				Field:    40.20,
				Operator: ">",
				Value:    20.30,
			},
			{
				Controls: Controls{
					Combinator: "and",
					Rules: []Rules{
						{
							Id:       "2e4fb517-4db7-46fc-90a7-542c8ec08f63",
							Field:    "Steve",
							Operator: "=",
							Value:    "Steve",
						},
					},
				},
			},
		},
	}

	ok, err := control.Control(data)

	if err != nil {
		t.Error(err)
		return
	}

	if !ok {
		t.Error("combinator AND returns false, test FAILED")
		return
	}

	t.Log("combinator AND returns true, test PASSED")
}

func TestOrCombinator(t *testing.T) {

	data := Controls{
		Id:         "0a6ccc76-0786-4bd5-99c2-3d516bca165a",
		Combinator: "or",
		Rules: []Rules{
			{
				Id:       "2e4fb517-4db7-46fc-90a7-542c8ec08f63",
				Field:    "Steve",
				Operator: "=",
				Value:    "Steve",
			},
			{
				Id:       "13a7bab3-c785-4403-8b2e-39c27cccc703",
				Field:    10.0,
				Operator: ">",
				Value:    20.50,
			},
			{
				Id:       "13a7bab3-c785-4403-8b2e-39c27cccc703",
				Field:    30.0,
				Operator: "<",
				Value:    20.30,
			},
			{
				Controls: Controls{
					Combinator: "or",
					Rules: []Rules{
						{
							Id:       "2e4fb517-4db7-46fc-90a7-542c8ec08f63",
							Field:    "Steve",
							Operator: "=",
							Value:    "Steve",
						},
					},
				},
			},
		},
	}

	ok, err := control.Control(data)

	if err != nil {
		t.Error(err)
	}

	if !ok {
		t.Error("combinator OR returns false, test FAILED")
	}

	t.Log("combinator OR returns true, test PASSED")
}
