sum_of_squares = 0
square_of_sum = 0
for i in range(1,101):
    sum_of_squares += i**2
    square_of_sum += i

square_of_sum *= square_of_sum

difference = square_of_sum - sum_of_squares

print difference
