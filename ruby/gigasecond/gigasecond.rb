class Gigasecond
  GIGASECOND = 10 ** 9
  def self.from(t)
    t + GIGASECOND
  end
end

module BookKeeping
  VERSION = 6
end
