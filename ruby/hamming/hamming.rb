class Hamming
  def self.compute(s1, s2)
    raise ArgumentError if s1.length != s2.length
    return 0 if s1 == s2
    a, b = s1.chars, s2.chars
    a.zip(b).count { |x, y| x != y }
  end
end

module BookKeeping
  VERSION = 3
end
