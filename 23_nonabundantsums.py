import numpy as np
import ast
from datetime import datetime
start_time = datetime.now()

# Find the sum of all the positive integers which cannot be written as the sum of two abundant number
def is_abundant(number):
    mysum = 1 # Can always divide by 1, so start looking at divisor 2
    for divisor in range(2, int(round(number / 2 + 1))):
        if number % divisor == 0:
            mysum += divisor
    if mysum > number:
        return True
    else:
        return False

# Find if number can be written as a sum of abundant numbers by subtracting by an abundant number and see if the difference is in the list of abundant numbers
"""
# Used to find the abundant numbers
# Comment this section during developement
abundant_numbers = []
for number in range(12, 28123 + 1):
    if is_abundant(number):
        abundant_numbers.append(number)
f = open('23_saved_abundant_numbers.txt','w')
f.write(str(abundant_numbers))
f.close()
print "Found all the abundant_numbers"
"""

f = open('23_saved_abundant_numbers.txt','r')
abundant_string = f.read()
f.close()
abundant_numbers = ast.literal_eval(abundant_string)

ans = 0
go_on = False
for number in range(1, 28123 + 1):
    for abundant in abundant_numbers:
        if abundant >= number:
            break
        elif (number - abundant) not in abundant_numbers:
                print "found {}".format(number)
                ans += number
                break

print "Sum of all positive integers which cannot be written as the sum of two abundant numbers is {}".format(ans)

print "This whole thing took {}".format(datetime.now() - start_time)
