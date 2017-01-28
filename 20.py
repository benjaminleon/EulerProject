import math

number = str(math.factorial(100))
print number

ans = 0
for digit in number:
    ans += int(digit)
    
print ans

