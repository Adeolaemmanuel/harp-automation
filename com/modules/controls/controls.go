package controls

import (
	"errors"
)

var iF = If{}

//Func called to control which combinator is to be called
func (con Controls) Control(payload Controls) (bool, error) {
	switch payload.Combinator {
	case AND:
		return con.AndCombinator(payload)
	case OR:
		return con.OrCombinator(payload)
	default:
		return false, errors.New("invalid combinator")
	}
}

//Func called for AND combinator
func (con Controls) AndCombinator(payload Controls) (bool, error) {

	var e error
	var b bool

	for _, r := range payload.Rules {
		data := mapName()

		ok, err := iF.IfControl(data[r.Operator], r.Field, r.Value)

		if len(r.Controls.Combinator) > 0 {
			return	con.Control(r.Controls)
		}

		if err != nil {
			e = err
			break
		}

		if !ok {
			b = ok
			break
		}

		b = ok
		
	}

	return b, e
}

//Func called for OR combinator process
func (con Controls) OrCombinator(payload Controls)  (bool, error) {
	var e error
	var b []bool

	for _, r := range payload.Rules {
		data := mapName()

		ok, err := iF.IfControl(data[r.Operator], r.Field, r.Value)

		if len(r.Controls.Combinator) > 0 {
			return	con.Control(r.Controls)
		}

		if err != nil {
			e = err
			break
		}
		
		b = append(b, ok)
	}

	return isContainsTrue(b), e
}

//Maps the operators key to each other
func mapName() map[string] string {

	data := make(map[string] string)

	data[IsEquals] = IsEquals
	data[IsLargerEqualNumber] = IsLargerEqualNumber
	data[IsLargerNumber] = IsLargerNumber
	data[IsLesserEqualNumber] = IsLesserEqualNumber
	data[IsLesserNumber] = IsLesserNumber
	data[IsNotEquals] = IsNotEquals
	data[IsStringContain] = IsStringContain
	data[IsStringDoesNotBeginWith] = IsStringDoesNotBeginWith
	data[IsStringDoesNotEndWith] = IsStringDoesNotEndWith
	data[IsStringEmpty] = IsStringEmpty
	data[IsStringEndWith] = IsStringEndWith
	data[IsStringNotContain] = IsStringNotContain
	data[IsStringStartWith] = IsStringStartWith

	return data
}

//Check if value true can be found in the array
func isContainsTrue(b []bool) bool {
	var bo bool

	for _, v := range b {
		if v  {
			bo = v
		}
	}

	return bo
}