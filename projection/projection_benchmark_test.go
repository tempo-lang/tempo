package projection_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tempo-lang/tempo/compiler"
)

func BenchmarkProjection(b *testing.B) {
	paths, err := filepath.Glob(filepath.Join("testdata", "examples", "*.tempo"))
	if err != nil {
		b.Fatal(err)
	}

	type Sample struct {
		name   string
		source string
	}

	samples := []Sample{}

	for _, path := range paths {
		dat, err := os.ReadFile(path)
		if err != nil {
			b.Fatal(err)
		}

		samples = append(samples, Sample{
			name:   filepath.Base(path),
			source: string(dat),
		})
	}

	for _, sample := range samples {
		b.Run(sample.name, func(b *testing.B) {
			for b.Loop() {
				input := antlr.NewInputStream(sample.source)
				_, errors := compiler.Compile(input, nil)
				if len(errors) > 0 {
					b.Fatalf("tempo compiler error: %v", errors[0])
				}
			}
		})
	}
}
