package provider

type OpenAiProvider struct {
	name string
}

func (p *OpenAiProvider) Name() string {
	return p.name
}
