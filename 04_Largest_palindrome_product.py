largest_palindrome = 0
for i in range(100,1000):
    for j in range(i,1000): # Don't try 2*3 and 3*2
        number = str(i * j)
        
        failure = False
        for idx in range(3):
            
            if number[idx] != number[-(idx+1)]:
                failure = True
                break
            
        if not failure and int(number) > largest_palindrome:
            largest_palindrome = int(number)
            print "updating", int(number)
          
print largest_palindrome
"""
1 1 1 1 1
1 2 3 4 5

2 2 2 2 2
  2 3 4 5

3 3 3 3 3
    3 4 5
"""
