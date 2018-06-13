class Complement
  def self.of_dna(dna)
    return '' if dna.match?(/[^GCTA]/)
    dna.tr('GCTA', 'CGAU')
  end
end

module BookKeeping
  VERSION = 4
end
