alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

with open('22_names.txt', 'r') as myfile:
    namelist = myfile.read().split(',')
    namelist.sort()

    namescore = 0
    for idx in range(len(namelist)):
        namecount = 0
        name = namelist[idx]
        lettercount = []

        for letter in name:
            if letter == "\"":
                continue
            print letter, alphabet.find(letter) + 1
            namecount += alphabet.find(letter) + 1
            lettercount.append((alphabet.find(letter)) + 1)
            
        namescore += (idx + 1) * namecount
        print (idx + 1), namelist[idx], namecount, lettercount

print namescore
