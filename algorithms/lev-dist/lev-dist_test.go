package levdist

import "testing"

var (
	words = []string{
		"a", "an", "ace", "act",
		"win", "zip", "zero", "alive", "above", "after", "again", "agree",
		"number", "oceanic",
		"tornado", "umbrella",
		"yesterday", "zodiac",
		"guitar", "history",
		"marathon", "notebook",
		"treasure", "vacation",
		"zookeeper", "alignment",
		"elevation", "foundation",
		"lightning", "meditation",
		"transformation",
		"championship",
		"inspiration",
		"organization",
		"verification",
	}
)

func TestLevDistSlice(t *testing.T) {
	a, b := "apple", "pears"
	expected := 5
	distance := levDistSlice(a, b)

	if distance != expected {
		t.Fatalf("Expected distance of %d for '%s' and '%s', but got %d\n", expected, a, b, distance)
	}

	a, b = "justify", "just"
	expected = 3
	distance = levDistSlice(a, b)

	if distance != expected {
		t.Fatalf("Expected distance of %d for '%s' and '%s', but got %d\n", expected, a, b, distance)
	}

	a, b = "execution", "intention"
	expected = 5
	distance = levDistSlice(a, b)

	if distance != expected {
		t.Fatalf("Expected distance of %d for '%s' and '%s', but got %d\n", expected, a, b, distance)
	}
}

func BenchmarkLevDistSlice(b *testing.B) {
	for i := range b.N {
		b.StopTimer()
		index := i % len(words)
		word1, word2 := words[index], words[len(words)-(index+1)]
		b.StartTimer()
		levDistSlice(word1, word2)
	}
}

func TestLevDistMatrix(t *testing.T) {
	a, b := "apple", "pears"
	expected := 5
	distance := levDistMatrix(a, b)

	if distance != expected {
		t.Fatalf("Expected distance of %d for '%s' and '%s', but got %d\n", expected, a, b, distance)
	}

	a, b = "justify", "just"
	expected = 3
	distance = levDistMatrix(a, b)

	if distance != expected {
		t.Fatalf("Expected distance of %d for '%s' and '%s', but got %d\n", expected, a, b, distance)
	}

	a, b = "execution", "intention"
	expected = 5
	distance = levDistMatrix(a, b)

	if distance != expected {
		t.Fatalf("Expected distance of %d for '%s' and '%s', but got %d\n", expected, a, b, distance)
	}
}

func BenchmarkLevDistMatrix(b *testing.B) {
	for i := range b.N {
		b.StopTimer()
		index := i % len(words)
		word1, word2 := words[index], words[len(words)-(index+1)]
		b.StartTimer()
		levDistMatrix(word1, word2)
	}
}

func TestLevDistTwoRows(t *testing.T) {
	a, b := "apple", "pears"
	expected := 5
	distance := levDistTwoRows(a, b)

	if distance != expected {
		t.Fatalf("Expected distance of %d for '%s' and '%s', but got %d\n", expected, a, b, distance)
	}

	a, b = "justify", "just"
	expected = 3
	distance = levDistTwoRows(a, b)

	if distance != expected {
		t.Fatalf("Expected distance of %d for '%s' and '%s', but got %d\n", expected, a, b, distance)
	}

	a, b = "execution", "intention"
	expected = 5
	distance = levDistTwoRows(a, b)

	if distance != expected {
		t.Fatalf("Expected distance of %d for '%s' and '%s', but got %d\n", expected, a, b, distance)
	}
}

func BenchmarkLevDistTwoRows(b *testing.B) {
	for i := range b.N {
		b.StopTimer()
		index := i % len(words)
		word1, word2 := words[index], words[len(words)-(index+1)]
		b.StartTimer()
		levDistTwoRows(word1, word2)
	}
}
