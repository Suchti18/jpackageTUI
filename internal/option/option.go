package option

type Option struct {
	optionName      string
	optionDesc      string
	optionCommand   string
	optionOs        Os
	optional        bool
	singleUse       bool
	possibleOptions []string
}

func NewOption(optionName string, optionDesc string, optionCommand string, optionOs Os, optional bool, singleUse bool,
	options []string) *Option {
	return &Option{
		optionName:      optionName,
		optionDesc:      optionDesc,
		optionCommand:   optionCommand,
		optionOs:        optionOs,
		optional:        optional,
		singleUse:       singleUse,
		possibleOptions: options,
	}
}

func (opt *Option) GetOptionName() string {
	return opt.optionName
}

func (opt *Option) GetOptionDesc() string {
	return opt.optionDesc
}

func (opt *Option) GetOptionCommand() string {
	return opt.optionCommand
}

func (opt *Option) GetOptionOs() Os {
	return opt.optionOs
}

func (opt *Option) IsOptional() bool {
	return opt.optional
}

func (opt *Option) IsSingleUse() bool {
	return opt.singleUse
}

func (opt *Option) GetPossibleOptions() []string {
	return opt.possibleOptions
}
