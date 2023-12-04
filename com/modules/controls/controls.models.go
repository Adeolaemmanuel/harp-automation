package controls

const (
	AND = "and"
	OR = "or"
	IF = "if"
	SWITCH = "switch"
)

const (
	IsEquals string = "="
	IsNotEquals string = "!="
	IsStringContain string = "contains"
	IsStringNotContain string = "doesNotContain"
	IsStringEndWith string = "endsWith"
	IsStringDoesNotBeginWith string = "doesNotBeginWith"
	IsStringDoesNotEndWith string = "doesNotEndWith"
	IsStringStartWith string = "beginsWith"
	IsStringEmpty string = "isStringEmpty"
	IsLargerNumber string = ">"
	IsLesserNumber string = "<"
	IsLesserEqualNumber string = "<="
	IsLargerEqualNumber string = ">="
)

type If struct {}

type Switch struct {}

type Controls struct {
	Id string `json:"id"`
	Combinator string `json:"combinator"`
	Rules []Rules`json:"rules"`
}

type Rules struct {
	Id string `json:"id"`
	Field any `json:"field"`
	Not bool `json:"not"`
	Value any `json:"value"`
	Operator string `json:"operator"`
	Controls
}

