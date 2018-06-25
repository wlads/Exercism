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
		go countFreqChan(s, c)
	}
	m := FreqMap{}
	for range ary {
		for letter, count := range <-c {
			m[letter] += count
		}
	}
	return m
}

func countFreqChan(s string, c chan FreqMap) {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	c <- m
}

func reduceMap(m1, m2 FreqMap) FreqMap {
	for k1 := range m1 {
		if _, ok := m2[k1]; ok {
			m1[k1] += m2[k1]
		}
	}
	return m1
}
