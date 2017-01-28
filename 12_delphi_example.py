import numpy as np
from datetime import datetime
startTime = datetime.now()
num = 0
j = 0

done = False
while not done:
    j += 1
    f = 0 
    num = num + j 
    for n in range(1, int(np.round(np.sqrt(num)))):
        if num % n == 0:
            f = f + 2 
    if f >= 500:
        done = True

print num
print datetime.now() - startTime
