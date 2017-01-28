from datetime import datetime
startTime = datetime.now()
end = 20
d = {(end,end): 0}

def calculate(x,y):

    if (x,y) in d:
        return d[(x,y)]
       
    elif x == y:
        ans = 2*calculate(x + 1,y)
        d[(x,y)] = ans
        return ans
    
    elif x < end and y < end:
        ans = calculate(x+1,y) + calculate(x,y+1)
        d[(x,y)] = ans
        return ans
    else:
        d[(x,y)] = 1
        return 1
        
print 'answer is:', calculate(0,0)
#for i in range(end):
#    print end - i, 'boxes:', d[(i,i)]

print datetime.now() - startTime
