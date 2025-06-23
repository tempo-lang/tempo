package projection

// SourceFile contains the entire source for a projected choreography from the perspective of all roles.
// It is used as the root under code generation.
type SourceFile struct {
	Interfaces []*ChoreographyInterface
	Structs    []*ChoreographyStruct
	Chors      []*Choreography
}

// NewSourceFile constructs a new source file.
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
