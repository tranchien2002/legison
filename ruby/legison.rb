require "./regex"


class Section
	attr_accessor :articles, :header

  def initialize header
    @articles = Array.new
		@header = header
  end


	def add_article article
		@articles.push(article)
	end
end

class Chapter
	attr_accessor :sections, :articles, :header


  def initialze header
		@header = header
		@sections = Array.new
		@articles = Array.new
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
  def initialize
    @base_ons = Array.new
    @signs = Array.new
    @chapters = Array.new
  end
end


section.add_article "điều 1"
p section.articles


# file = File.new("test.txt","r+")
# file.each_line do |line|
# 	puts line
# 	puts "88888888888888888888888888"
# end

# test = Legison.new
