package schema

type (
	inputType int
)

const (
	NumberInput inputType = iota + 1
	StringInput
	DateInput
	OptionInput
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
	InputValue               string
	InputErr                 error
	InputOptions             []*InputOption // Optional
	InputOptionValueSelected string         // Optional
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

func NewInput(
	inputTitle string,
	inputName string,
	inputType inputType,
	inputValue string,
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
	}
}
