class Hamming
  def self.compute(s1, s2)
    raise ArgumentError if s1.length != s2.length
    return 0 if s1 == s2
    index = 0
    hamming_distance = 0
    s1.each_char do |c|
      if c != s2[index]
        hamming_distance += 1
      end
      index += 1
    end
    return hamming_distance
  end
end

module BookKeeping
  VERSION = 3
end
