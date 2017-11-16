require "./regex"
require "pry"
class Topic

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

	attr_accessor :topics, :articles, :header


  def initialize header
		@header = header
		@topics = Array.new
		@articles = Array.new

  end

	def add_topic topic
		@topics.push(topic)
	end

	def add_article article
		@articles.push(article)
	end
end

class Legison
	attr_accessor :sigs, :baseons, :chapters, :header

	def initialize header
		@header = header
    @baseons = Array.new
    @signs = Array.new
    @chapters = Array.new
	end
	
	def add_baseons baseon
		@baseons.push(baseon)
	end
end
topic_obj = Array.new
chapter_obj = Array.new

string = File.open("test.txt","r+").read

chapters = []
string.scan(/\n(Chương|CHƯƠNG)\s(M{1,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})|M{0,4}(CM|C?D|D?C{1,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})|M{0,4}(CM|CD|D?C{0,3})(XC|X?L|L?X{1,3})(IX|IV|V?I{0,3})|M{0,4}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|I?V|V?I{1,3}))\s+(.+)|(\.\/\.)/)  do |e|
	chapters << [e[14], $~.offset(0)]
  obj = Chapter.new(e[14])
  chapter_obj.push(obj)
end


topics = []
string.scan(/\n(Mục|MỤC)\s*(\d+)\s*(.+)/) do |e|
	topics << [e[2], $~.offset(0)]
  obj = Topic.new(e[2])
  topic_obj.push(obj)
end

articles = []
string.scan(/(\n.iê?.(u|ù)\s+(\d+)\.\s+(.+))|(\.\/\.)/) do |e|
	articles << [e[3], $~.offset(0)]
end

# p topic_obj[0].add_article(articles[0][0])
# topic_obj[0].add_article("adsf")
# p topic_obj[0]
temp = Array.new
count = 0

articles.each do |a|
  p count
  if( (a[1][0] >  topics[count][1][0]) && (a[1][0] < topics[count][1][1]) )

    p a[1][0]
    # topic_obj[0].add_article(a[0])
    # p topic_obj[0]
    # articles.delete(a)

  elsif( a[1][0] < topics[count][1][0] )
    next
  else
    count += 1
  end
end