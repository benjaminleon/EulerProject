number = 600851475143.0
possible = [2] + range(3,10000,2)

try:
    for current in range(len(possible)):
        for other in range(current):
            if possible[current] % possible[other] == 0:
                del possible[current]
except IndexError: # Too long range 'cus we're deleting elements
    pass

primes = possible
#print primes

largest_prime = 1
for i in range(len(primes)):    
    prime = primes[i]

    if prime > number:
        break
    
    while number % prime == 0:
        largest_prime = prime
        number = number / prime

print largest_prime
        
