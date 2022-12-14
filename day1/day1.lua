
-- http://lua-users.org/wiki/FileInputOutput

-- see if the file exists
function file_exists(file)
  local f = io.open(file, "rb")
  if f then f:close() end
  return f ~= nil
end

-- get all lines from a file, returns an empty
-- list/table if the file does not exist
function lines_from(file)
  if not file_exists(file) then return {} end
  local lines = {}
  for line in io.lines(file) do
    lines[#lines + 1] = line
  end
  return lines
end

-- tests the functions above
local file = 'input.txt'
local lines = lines_from(file)

-- print all line numbers and their contents
Counts = {}
local count = 0
for k,v in pairs(lines) do
  if v ~= "" then
    count = count + v
  else
    table.insert(Counts,count)
    count = 0
  end
end
table.sort(Counts)
print("max",Counts[#Counts])
print("top three",Counts[#Counts]+Counts[#Counts-1]+Counts[#Counts-2])
