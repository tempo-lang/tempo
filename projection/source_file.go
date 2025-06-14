package projection

type SourceFile struct {
	Interfaces []*ChoreographyInterface
	Structs    []*ChoreographyStruct
	Chors      []*Choreography
}

func NewSourceFile() *SourceFile {
	return &SourceFile{
		Chors: []*Choreography{},
	}
}

func (s *SourceFile) AddChoreography(chor *Choreography) {
	s.Chors = append(s.Chors, chor)
}

func (s *SourceFile) AddInterface(st *ChoreographyInterface) {
	s.Interfaces = append(s.Interfaces, st)
}

func (s *SourceFile) AddStruct(st *ChoreographyStruct) {
	s.Structs = append(s.Structs, st)
}
