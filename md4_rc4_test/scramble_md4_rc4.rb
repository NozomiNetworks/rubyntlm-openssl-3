#!ruby

require_relative '../lib/net/ntlm/md4.rb'
require_relative '../lib/net/ntlm/rc4.rb'

def md4(data)
  Net::NTLM::Md4.digest(data)
end

def rc4(key, data)
  Net::NTLM::Rc4.new(key).encrypt(data)
end

def read_all(f, size)
  res = f.read(size)
  return nil if res && res.size < size
  res
end

def process(name)
  f = File.open(name)
  result = md4('')
  loop {
    sizes = read_all(f, 2)
    break unless sizes
    key = read_all(f, sizes.bytes[0] + 1)
    break unless key
    value = read_all(f, sizes.bytes[1] + 1)
    break unless value
    result = md4(rc4(key,value) + result)
  }
  f.close
  result
end

ARGV.each do |name| 
  puts "#{process(name).unpack1("H*")}  #{name}"
end
