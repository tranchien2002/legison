require "./regex"


class Section
	attr_accessor :articles

  def initialize
    @articles = Array.new
  end


	def add_article article
		@articles.push(article)
	end
end

class Chapter
	attr_accessor :sections, :articles
	@sections = Array.new
	@articles = Array.new

  def initialze
  end

	def add_section section
		@sections.push(section)
	end

	def add_article article
		@articles.push(article)
	end
end

class Legison
	attr_accessor :sigs, :base_ons, :chapters
	@signs
	@base_ons = Array.new
	@chapters
  def initialize

  end
end

section = Section.new
section.add_article "điều 1"
p section.articles


# file = File.new("test.txt","r+")
# file.each_line do |line|
# 	puts line
# 	puts "88888888888888888888888888"
# end

# test = Legison.new
