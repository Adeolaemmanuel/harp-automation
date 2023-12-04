package controls

import (
	"testing"
)


var control = Controls{}

func TestAndCombinator(t *testing.T) {

	data := Controls{
		Id: "0a6ccc76-0786-4bd5-99c2-3d516bca165a",
		Combinator: "and",
		Rules: []Rules{
			{
				Id: "2e4fb517-4db7-46fc-90a7-542c8ec08f63",
				Field: "Steve",
				Operator: "=",
				Value: "Steve",
			},
			{
				Id: "13a7bab3-c785-4403-8b2e-39c27cccc703",
				Field: 40,
				Operator: ">",
				Value: 20,
			},
			{
				Id: "13a7bab3-c785-4403-8b2e-39c27cccc703",
				Field: 40,
				Operator: ">",
				Value: 20,
			},
			{
				Controls: Controls{
					Combinator: "and",
					Rules: []Rules{
						{
							Id: "2e4fb517-4db7-46fc-90a7-542c8ec08f63",
							Field: "Steve",
							Operator: "=",
							Value: "Steve",
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
		t.Error("combinator AND returns false, test failed")
	}

	t.Log(ok)
}

func TestOrCombinator(t *testing.T) {

	data := Controls{
		Id: "0a6ccc76-0786-4bd5-99c2-3d516bca165a",
		Combinator: "and",
		Rules: []Rules{
			{
				Id: "2e4fb517-4db7-46fc-90a7-542c8ec08f63",
				Field: "Steve",
				Operator: "=",
				Value: "Steve",
			},
			{
				Id: "13a7bab3-c785-4403-8b2e-39c27cccc703",
				Field: 30,
				Operator: ">",
				Value: 20,
			},
			{
				Id: "13a7bab3-c785-4403-8b2e-39c27cccc703",
				Field: 30,
				Operator: ">",
				Value: 20,
			},
			{
				Controls: Controls{
					Combinator: "and",
					Rules: []Rules{
						{
							Id: "2e4fb517-4db7-46fc-90a7-542c8ec08f63",
							Field: "Steve",
							Operator: "=",
							Value: "Steve",
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
		t.Error("combinator OR returns false, test failed")
	}

	t.Log(ok)
}