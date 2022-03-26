package device

func NewSetter(commandSetter CommandSetter, response Response) *setter {
	return &setter{CommandSetter: commandSetter, Response: response}
}

type setter struct {
	CommandSetter
	Response
}

func (s setter) GetResponseLength() int {
	return s.Response.GetLength()
}
