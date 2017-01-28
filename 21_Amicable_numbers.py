# Function which finds sum of divisors of a given number
def sum_divisors(number):
    mysum = 0
    if number < 4:
        return 1
    for divisor in range(1, int(round(number / 2 + 1))):
        if number % divisor == 0:
            mysum += divisor
    return mysum

# First pass, find the sum of the divisors for all numbers
max_number = 10000
list_of_sums = [0] * max_number
for i in range(max_number):
    list_of_sums[i] = sum_divisors(i)

# Second pass, find the amicable numbers
amicable = []
for idx in range(1, max_number):
    divisor_sum = list_of_sums[idx]

    if divisor_sum >= max_number:
        continue

    if list_of_sums[divisor_sum] == idx and idx != divisor_sum:
        amicable.append(idx)
        amicable.append(divisor_sum)
        print "Amicable pair found!", idx, "and", divisor_sum
        
print "Sum of the amicable numbers, where no number is larger than", max_number, "is", sum(set(amicable))
