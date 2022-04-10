import decimal
import re
r = re.compile(r"(.+?)\1+")

longest = 0
winning_idx = 0
for i in range(1, 1000):

    m = r.findall(repr(float(1.0) / float(i)))
    
    #print (1.0 / i), r.findall(repr(1.0 / i)), "index:", i

    if m:
        cur_len = len(max(r.findall(repr(float(1.0) / float(i)))))
    else:
        continue

    if cur_len > longest:
        longest = cur_len
        winning_idx = i

print "Answer is index:", winning_idx, longest, "and pattern", r.findall(repr(1.0 / winning_idx))

print('%.16f' % (float(1.0) / float(winning_idx)))
