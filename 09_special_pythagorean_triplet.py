import numpy as np

maxnumber = int(1000.0/3)

done = False
for a in range(1, maxnumber):
    if done:
        break
    
    for b in range(a, 1000):
        c = 1000 - a - b

        if a**2 + b**2 == c**2:
            print 'a, b, c:', a, b, c
            print 'a^2, b^2, c^2:', a**2, b**2, c**2
            print 'a + b + c:', a + b + c
            print 'a*b*c:', a*b*c
            done = True
            break


