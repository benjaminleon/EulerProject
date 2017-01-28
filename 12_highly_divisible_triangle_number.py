import math
import numpy as np
from datetime import datetime
startTime = datetime.now()
natural = range(1, 500000)

# Get list of prime numbers
def prime_list(max_prime):
    primes = range(2, max_prime)
    length = len(primes)

    for idx in range(len(primes)):
        p = primes[idx]
        if p == 0:
            continue
        
        # No multiples of any prime is a prime
        for i in range(2, (length + 1) / p + 1):
            primes[p*i - 2] = 0
    
    primes = [y for y in primes if y != 0]
    return primes

# Construct list of triangles (= cummulative sum)
triangles = np.zeros(len(natural)).astype(np.int)
triangles[0] = 1
for i in range(1,len(natural)):
    triangles[i] = natural[i] + triangles[i - 1]

# Find list of prime numbers
primes = prime_list(int(np.sqrt(triangles[-1]))) # Only need this many primes
    
done = False
for triangle_idx in range(len(triangles)):

    if done:
        break
    tri = float(triangles[triangle_idx])    

    # Remove primes which does not constitute the considered number
    lego = [prime for prime in  primes if tri % prime == 0]

    new_divisors = list(lego)
    stored = []
    new_found = True
    while new_found:

        # Fill with all combinations of primes and their products
        tmp = np.zeros(len(lego)*len(new_divisors)).astype(np.int)

        for i in range(len(lego)):
            for j in range(len(new_divisors)):
                
                # Make all combinations
                tmp[i*len(new_divisors) + j] = lego[i]*new_divisors[j]

        tmp2 = [new for new in tmp if tri % new == 0]
        if set(new_divisors) == set(tmp2) or len(tmp2) == 0:
            new_found = False
        else:
            stored += new_divisors
            new_divisors = list(set(tmp2))

    ans = len(stored) + 1 # Itself

    if ans >= 500: # Don't try more triangle values
        done = True 

print 'triangle value', int(tri), 'with index', triangle_idx, 'gives', ans, 'possible divisors'
print 'primes:', lego
#print 'Possible divisors:', sorted(stored)
print datetime.now() - startTime

"""    
    prime_dict = {}

    for p in range(len(lego)):
        prime_dict[lego[p]] = 0
        
    # Lego are the unique primes which builds the number.
    # Find out how many primes the number is made from
    nr_of_factors = 0
    tmp_tri = tri
    
    for i in range(len(lego)):
        while tmp_tri % lego[i] == 0:
            tmp_tri /= lego[i]
            prime_dict[lego[i]] += 1
            nr_of_factors += 1

    print 'tri:', tri
    print 'prime_dict', prime_dict
"""


"""
        # When chosing 2 primes to make a factor for the number, and
        # the number is made from, let's say 3 of the same prime, then
        # limit those primes to 2 so the 'a chose b' doesn't produce
        # identical copies. Chosing 2 out of [5, 5, 5] should only give
        # [5, 5], i.e. there is only one way to do it.
        chose_from = np.sum([min(prime_dict[lego[x]], i) for x in range(len(prime_dict))])
        print 'chose', i, 'from', chose_from,':',math.factorial( chose_from ) / (math.factorial( chose_from - i ) * math.factorial( i ))
        ans += math.factorial( chose_from ) / (math.factorial( chose_from - i ) * math.factorial( i ))
"""


"""
    # With tri as 360, prime_dict is {2: 3, 3:2, 5:1}

    # When grabbing 2 legos, we can take 0,1 or 2 of 2,
    # 0,1 or 2 from 3, and 0 or 1 from 5.

    # When grabbing 3 legos, we can take 0,1,2 or 3 of 2,
    # 0,1 or 2 from 3, and 0 or 1 from 5.
    # Search for these combinations where the sum of the
    # number of lego pieces are 3.

    # When grabbing 4 legos, we have the same options, but
    # now we search for combinations where the sum is 4

    # This generalizes to that we can take values from
    # a range from 0 to min(#of pieces, #of legos in bin)
    # in every bin.

    # (Start searching from the bin with fewest legos to
    # terminate search early.)
    
    ans = 1 # Instead of reaching nr_of_factors which will give 1
    
    for i in range(1, nr_of_factors): # Pick 1,2,3...

        select = []
        for piece_idx in range(len(lego)):
            piece = lego[piece_idx]

            # From 2*2*2, we can take 0,1,2 or 3 2's
            select.append(range(prime_dict[piece] + 1) )
        print select
        print len(select)
        print select[0][:]

        for piece_idx in range(len(lego)):
            hej = select[piece_idx][i] + select[piece_idx]
        
        tjubadoo
"""            
