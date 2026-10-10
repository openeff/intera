package provider

type LMStudioProvider struct {
	name string
}

func (p *LMStudioProvider) Name() string {
	return p.name
}
