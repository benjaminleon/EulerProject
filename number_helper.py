def is_abundant(number):
    mysum = 1 # Can always divide by 1, so start looking at divisor 2
    for divisor in range(2, int(round(number / 2 + 1))):
        if number % divisor == 0:
            mysum += divisor
    if mysum > number:
        return True
    else:
        return False
