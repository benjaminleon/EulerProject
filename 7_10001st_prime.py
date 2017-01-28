import numpy as np
primes = [2] + range(3,110000,2)
deleted = 0

try:
    for current in range(len(primes)):
        current -= deleted

        other = 0
        while True:
            # Don't need to try all primes
            if primes[other] > np.sqrt(primes[current]):
                break

            if primes[current] % primes[other] == 0:                
                del primes[current]
                deleted += 1
            other += 1
            
except IndexError: # Too long range 'cus we're deleting elements
    pass

print len(primes)
print primes[-1]

