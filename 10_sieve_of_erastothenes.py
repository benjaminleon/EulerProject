import numpy as np
from datetime import datetime
startTime = datetime.now()

max_prime = 2000000
primes = range(2,max_prime)
length = len(primes)

for idx in range(len(primes)):
    p = primes[idx]
    if p == 0:
        continue
        
    # No multiples of any prime is a prime
    for i in range(2, (length + 1) / p + 1):
        primes[p*i - 2] = 0
    
primes = [y for y in primes if y != 0]

print "Time to make list", datetime.now() - startTime

print 'sum_of_primes less than', max_prime,':', np.sum(primes)
