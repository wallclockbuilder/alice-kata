dict = Hash.new(0)
File.foreach('alice-in-wonderland.txt'){|line|
  line.scan(/\w+/){|word|
    dict[word.downcase] += 1
  }
}
p dict.sort_by { |k, v| v }.reverse[0..6]



#readline by line
#convert lines to words
#insert words into dictionary with count
#sort dictionary by value into array
#return top 7
