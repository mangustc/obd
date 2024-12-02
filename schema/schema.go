package schema

type (
	inputType int
)

const (
	NumberInput inputType = iota + 1
	StringInput
	DateInput
	OptionInput
	BooleanInput
)

type TableHeaderColumn struct {
	Name    string
	Percent int
}

type InputOption struct {
	InputOptionLabel string
	InputOptionValue string
}

type Input struct {
	InputTitle               string // Optional
	InputName                string
	InputType                string
	InputValue               any
	InputErr                 error
	InputOptions             []*InputOption // Optional
	InputOptionValueSelected string         // Optional
	InputEditable            bool           // Optional
}

func NewTableHeaderColumn(name string, percent int) *TableHeaderColumn {
	if name == "" {
		panic("Header name can't be zero")
	}
	return &TableHeaderColumn{
		Name:    name,
		Percent: percent,
	}
}

func setNotEditableInput(input *Input) {
	input.InputEditable = false
}

func NewInputNotEditable(inputValue any) *Input {
	switch inputValue.(type) {
	case int:
		input := NewInput("", ".", NumberInput, inputValue, nil, nil, "")
		setNotEditableInput(input)
		return input
	case string:
		input := NewInput("", ".", StringInput, inputValue, nil, nil, "")
		setNotEditableInput(input)
		return input
	case bool:
		input := NewInput("", ".", BooleanInput, inputValue, nil, nil, "")
		setNotEditableInput(input)
		return input
	default:
		panic("InputValue is not set to a supported value type")
	}
}

func NewInput(
	inputTitle string,
	inputName string,
	inputType inputType,
	inputValue any,
	inputErr error,
	inputOptions []*InputOption,
	inputOptionValueSelected string,
) *Input {
	if inputName == "" {
		panic("One or more neccesary arguments are zero")
	}

	var inputTypeStr string
	switch inputType {
	case NumberInput:
		inputTypeStr = "number"
	case StringInput:
		inputTypeStr = "string"
	case DateInput:
		inputTypeStr = "date"
	case BooleanInput:
		inputTypeStr = "checkbox"
	case OptionInput:
		if inputOptions == nil {
			panic("Options arrray can not be nil")
		}
		for _, inputOption := range inputOptions {
			if inputOption == nil {
				panic("Input option can not be nil")
			}
			// NOTE: this function does not check for right for not repeating input options
		}
	default:
		panic("InputType is not set to a real inputType value")
	}

	return &Input{
		InputTitle:               inputTitle,
		InputName:                inputName,
		InputType:                inputTypeStr,
		InputValue:               inputValue,
		InputErr:                 inputErr,
		InputOptions:             inputOptions,
		InputOptionValueSelected: inputOptionValueSelected,
		InputEditable:            true,
	}
}
