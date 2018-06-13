class Squares
  def initialize(number)
    @number = number
  end
  def square_of_sum
    (1..@number).sum ** 2
  end
  def sum_of_squares
    (1..@number).map { |n| n ** 2 }.sum
  end
  def difference
    square_of_sum - sum_of_squares
  end
end

module BookKeeping
  VERSION = 4
end
