nums = []
freq = {}
with open('input.txt', 'r') as f:
    lines = f.readlines()
    for line in lines:
        if len(line) == 0:
            break
        a, b = map(int, line.split())
        nums.append(a)
        if b not in freq:
            freq[b] = 0
        freq[b] += 1
sum = 0
for i in nums:
    if i in freq:
        sum += i * freq[i]
print(sum)
