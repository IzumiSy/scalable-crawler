require "open-uri"
require "nokogiri"

crawling_url = ARGV[0]

unless crawling_url
  puts "URL to crawl is empty"
  exit 1
end

page = Nokogiri.HTML(open(crawling_url))

puts page.title
