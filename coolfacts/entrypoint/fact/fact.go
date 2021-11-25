package fact

type Fact struct {
	Image string `json:"image"`
	Description string `json:"description"`
}

type repo struct {
	facts []Fact
}

func (r *repo) getAll() []Fact{
	return r.facts
}

func (r *repo) add(f Fact){
	 r.facts = append(r.facts, f)
}