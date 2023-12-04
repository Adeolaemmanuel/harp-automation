package automation

type Automation struct {}

type AutomationNode struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Categories []string `json:"categories"`
	Credentials []AutomationCredential `json:"credentials"`
	RequestDefaults AutomationRequestDefault `json:"request_defaults"`
	RequestOptions string `json:"request_options"`
	Entities []AutomationEntity `json:"entities"`
	Operations []AutomationOperation `json:"operations"`
	Fields []AutomationField `json:"fields"`
}

type AutomationCredential struct {
	Data map[string]string `json:"data"`
	Name string `json:"name"`
	Required bool `json:"required"`
	Type string `json:"type"`
}

type AutomationRequestDefault struct {
	BaseURL string `json:"baseURL"`
	Headers string `json:"headers"`
	Description string `json:"description"`
}

type AutomationOperation struct {
	Key string `json:"key"`
	Label string `json:"label"`
	Description string `json:"description"`
	Type string `json:"type"`
	Fields []AutomationField
	Routing AutomationRouting `json:"routing"`
}

type AutomationField struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	Description string `json:"description"`
	Type string `json:"type"`
	Type_Options string `json:"type_options"`
	Display_Options string `json:"display_options"`
	Required bool `json:"required"`
	Is_Input bool `json:"is_input"`
}

type AutomationRouting struct {
	Method string `json:"method"`
	Url string `json:"url"`
}

type AutomationEntity struct {
	Key string `json:"key"`
	Label string `json:"label"`
	Description string `json:"description"`
	Operation AutomationOperation `json:"operation"`
}