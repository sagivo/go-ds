arr = Array.new(100000000) { |i| 0  }
File.readlines('arr.txt').each_with_index do |line,i|
  arr[i] = line.to_i #if i < arr.size and line.size > 0
end

start = Time.now
arr.sort
p Time.now - start