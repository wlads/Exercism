class Pangram
  def self.pangram?(phrase)
    phrase.downcase!
    ('a'..'z').to_a.all? { |c| phrase.include? c }
  end
end

module BookKeeping
  VERSION = 6
end
