def num_2_word_10_down(number):
    if number == 0:
        return 'zero'
    elif number == 1:
        return 'one'
    elif number == 2:
        return 'two'
    elif number == 3:
        return 'three'
    elif number == 4:
        return 'four'
    elif number == 5:
        return 'five'
    elif number == 6:
        return 'six'
    elif number == 7:
        return 'seven'
    elif number == 8:
        return 'eight'
    elif number == 9:
        return 'nine'

def ten_to_20(number):
    if number == 10:
        return 'ten'
    elif number == 11:
        return 'eleven'
    elif number == 12:
        return 'twelve'
    elif number == 13:
        return 'thirteen'
    elif number == 14:
        return 'fourteen'
    elif number == 15:
        return 'fifteen'
    elif number == 16:
        return 'sixteen'
    elif number == 17:
        return 'seventeen'
    elif number == 18:
        return 'eighteen'
    elif number == 19:
        return 'nineteen'

def num_2_word_twenty_up(number):
    if number == 2:
        return 'twenty'
    elif number == 3:
        return 'thirty' 
    elif number == 4:
        return 'forty'
    elif number == 5:
        return 'fifty'
    elif number == 6:
        return 'sixty'
    elif number == 7:
        return 'seventy'
    elif number == 8:
        return 'eighty'
    elif number == 9:
        return 'ninety'


totalcount = 0
for number in range(1,1000):
    count = 0
    numstr = ''
    n = number

    if number >= 100: # four hundred and fifty two comes in
        hundreds = number / 100 # take the 4 from 452
        numstr += num_2_word_10_down(hundreds) # add 'four' to the string

        if number - hundreds*100 == 0: # if we had 400 instead of 452
            numstr += 'hundred' # add 'hundred'
            count = len(numstr)
            totalcount += count
            print 'Number:', n, numstr, '\t', 'charcount:', count
            continue
        
        else:
            number -= hundreds*100 # make 452 become 52
            numstr += 'hundredand' # add 'hundredand', later, 'fiftytwo'
    
    if number >= 20:
        dec = number / 10
        numstr += num_2_word_twenty_up(dec) # fourty...
        if number - dec*10 != 0:
            numstr += num_2_word_10_down(number - dec*10) # ...five

    elif number >= 10:
        numstr += ten_to_20(number)

    else:
        numstr += num_2_word_10_down(number)

    count = len(numstr)

    totalcount += count
    
    print 'Number:', n, numstr, '\t', 'charcount:', count

print 'totalcount:', totalcount + len('onethousand')

