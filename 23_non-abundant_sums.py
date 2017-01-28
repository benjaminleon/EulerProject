import numpy as np
import ast
from number_helper import is_abundant # My own module
from datetime import datetime
start_time = datetime.now()

# Find if number can be written as a sum of abundant numbers by subtracting by an abundant number and see if the difference is in the list of abundant numbers

# Used to find the abundant numbers

abundant_numbers = []
for number in range(12, 28123 + 1):
    if is_abundant(number):
        abundant_numbers.append(number)

print "Found all the abundant_numbers"

""" Saves and reads abundant numbers
f = open('23_saved_abundant_numbers.txt','w')
f.write(str(abundant_numbers))
f.close()

f = open('23_saved_abundant_numbers.txt','r')
abundant_string = f.read()
f.close()
abundant_numbers = ast.literal_eval(abundant_string)
"""

def go_through_all(abundant_numbers):
    ans = 0
    go_on = False
    for number in range(1, 28123 + 1):
        if number % 500 == 0:
            print number

        num_found = False
        for abundant in abundant_numbers:    

            if abundant > number / 2: # Symmetry
                num_found = True
                break

            if (number - abundant) in abundant_numbers:
                break

        if num_found:
            ans += number
    return ans

def make_combinations(abundant_numbers):
    summable_numbers = []
    for idx, abundant1 in enumerate(abundant_numbers):
        for abundant2 in abundant_numbers[idx:]:
            if abundant1 + abundant2 < 28123 + 1:
                summable_numbers.append(abundant1 + abundant2)

    print "Found all summable numbers"

    #print "Got {} duplicates".format(len(summable_numbers) - len(list(set(summable_numbers)))) # Got 24205224 duplicates
                                     
    summable_numbers = list(set(summable_numbers))
    non_summable = range(1, 28123 + 1)
    for summable in summable_numbers:
        non_summable.remove(summable)
    return sum(non_summable)

ans = make_combinations(abundant_numbers)
    
print "Sum of all positive integers which cannot be written as the sum of two abundant numbers is {}".format(ans)

print "This whole thing took {}".format(datetime.now() - start_time)
# 9 minutes when reading the abundant numbers from file, using go_through_all
# 21 seconds without reading saved numbers, using make_combinations
