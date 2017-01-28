
sunday = 0
monday = 1

january = 0
february = 1
mars = 2
april = 3
may = 4
june = 5
july = 6
august = 7
september = 8
oktober = 9
november = 10
december = 11

day = monday
ans = 0

counting = False

def first_sunday(day, counting):
    if day == 0 and counting:
        return 1
    else:
        return 0

for year in range(1900, 2001):
    for month in range(12):

        print 'year, month:', year, month
        
        if year == 1901:
            counting = True
        
        if month == april or month == june or month == september or month == november:
            day = (day + 30) % 7
            
            ans += first_sunday(day, counting)

        elif month == february:
            if year % 4 == 0 and (year % 100 != 0 or year % 400 == 0):
                leap = 1
            else:
                leap = 0
                
            day = (day + 28 + leap) % 7
            ans += first_sunday(day, counting)

        else:
            day = (day + 31) % 7
            ans += first_sunday(day, counting)
                
print ans
