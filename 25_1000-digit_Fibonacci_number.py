f1 = 1
f2 = 1

nr_of_digits = 1000
counter = 1
while True:
    f1 = f1 + f2
    f2 = f1 + f2
    counter += 1
    if len(str(f2)) >= nr_of_digits:
        if len(str(f1)) >= nr_of_digits:
            print "Index {}".format(counter * 2 - 1)
            break
        else:
            print "Index {}".format(counter * 2)
            break
