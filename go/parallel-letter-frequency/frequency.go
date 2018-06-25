package letter

// FreqMap represents the frequency of letters
type FreqMap map[rune]int

// Frequency counts frequency of letters in a string
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

// ConcurrentFrequency counts the frequency of letters in a array of strings concurrently
func ConcurrentFrequency(ary []string) FreqMap {
	c := make(chan FreqMap)
	for _, s := range ary {
		go func(text string) { c <- Frequency(text) }(s)
	}
	m := FreqMap{}
	for range ary {
		for letter, count := range <-c {
			m[letter] += count
		}
	}
	return m
}

// not in use
func reduceMap(m1, m2 FreqMap) FreqMap {
	for k1 := range m1 {
		if _, ok := m2[k1]; ok {
			m1[k1] += m2[k1]
		}
	}
	return m1
}

// alternative sol sharing data structure: http://exercism.io/submissions/3cabd34a00324ba08d15a987ebc33395
// nice solution with wait group: http://exercism.io/submissions/c859d0df79204b148245a00de7fe60ff
