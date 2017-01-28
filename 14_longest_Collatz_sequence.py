from datetime import datetime
startTime = datetime.now()

longest_chain = 0
best_starting_number = 0

for i in range(1, 1000000, 2):
    n = i
    count = 1

    # Trace until 1
    while n != 1:
        count += 1
        if n % 2 == 0:
            n /= 2
        else:
            n = 3*n + 1
    #print 'i:', i, 'count:', count

    # Keep the best found so far
    if count > longest_chain:
        longest_chain = count
        best_starting_number = i
        
print 'winning number is:', best_starting_number
print 'It took ',datetime.now() - startTime    
