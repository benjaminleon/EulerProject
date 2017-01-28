from datetime import datetime
startTime = datetime.now()
maxmoves = 15

def calculate(x,y):

    if x == y == maxmoves:
        return 1
       
    elif x == y:
        return 2*calculate(x + 1,y)
    
    elif x < maxmoves and y < maxmoves:
        return calculate(x+1,y) + calculate(x,y+1)
    else:
        return 1
        
print calculate(0,0)
print datetime.now() - startTime
