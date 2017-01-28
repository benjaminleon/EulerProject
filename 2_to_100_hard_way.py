import numpy as np
def mult(first, second):
    print "multiplying:\n", second, "with\n", first

    str1 = str(first)
    str2 = str(second)
    print "i","j", ':', "product\n"

    total_prod = np.zeros(2*max(len(str1),len(str2)))
    
    for i in range(len(str1)):
        decimal = 0
        num1 = int(str1[-1-i])

        for j in range(len(str2)):
            num2 = int(str2[-1-j])
            prod = num1*num2 + decimal

            if prod + total_prod[-1-i-j] >= 10:
                decimal = int(str(prod + total_prod[-1-i-j])[0])
                current = prod - 10*decimal
            else:
                decimal = 0
                current = prod

            total_prod[-1-i-j] += current

        total_prod[-2-i-j] += decimal
        print total_prod,"\n"

    digsum = 0
    for i in range(len(total_prod)):
        digsum += total_prod[i]

    return int(digsum)

first = 85754
second = 459
print mult(first,second)
print first*second

