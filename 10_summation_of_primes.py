import numpy as np
from datetime import datetime

max_prime = 100
startTime = datetime.now()
primes = [2] + range(3,max_prime,2)

for current in range(len(primes)):
    if primes[current] == 1: # Not a prime
        continue
    
    other = 0
    while True:
        if primes[other] == 1: # Not a prime
            other += 1
            continue
        
        # Don't need to try all primes
        if primes[other] > np.sqrt(primes[current]):
            break
        
        if primes[current] % primes[other] == 0:                
            primes[current] = 1 # Not a prime in the list
            
        other += 1
        
# Remove the 1's
primes = [y for y in primes if y != 1]

print "Time to make list", datetime.now() - startTime

sum_of_primes = 0
for i in range(len(primes)):
    if primes[i] < 2000000:
        sum_of_primes += primes[i]

print 'number of primes:', len(primes)
print 'largest prime', primes[-1]
print 'sum of primes less than', max_prime, ':', sum_of_primes
