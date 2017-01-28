a = 1
b = 2

ans = 2

while True:
    a = a + b
    b = a + b
    if a >= 4000000 or b >= 4000000:
        break
    if a % 2 == 0:
        ans += a
    if b % 2 == 0:
        ans += b

print ans
        
