class Raindrops
  def self.convert(number)
    conv_num = ""
    conv_num << "Pling" if number % 3 == 0
    conv_num << "Plang" if number % 5 == 0
    conv_num << "Plong" if number % 7 == 0
    conv_num.empty? ? number.to_s : conv_num
  end
end

module BookKeeping
  VERSION = 3
end
