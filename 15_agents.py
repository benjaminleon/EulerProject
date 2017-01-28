maxmoves = 12
agents = []
count = 0

class Agent(object):

    def __init__(self, (x,y)):
        self.x = x
        self.y = y
        self.act()
        
    def duplicate(self):
        if self.x < maxmoves and self.y < maxmoves:
            # Spawn a new agent when there are 2 posisble moves
            newagent = Agent([self.x, self.y + 1])

    def move(self):
        if self.x < maxmoves:
            self.x += 1
        elif self.y < maxmoves:
            self.y += 1
        
    def act(self):
        done = False
        while not done:
            self.duplicate()
            self.move()
            if self.x == maxmoves and self.y == maxmoves:
                global count
                count += 1
                done = True


newagent = Agent([0,0])
print 'count:', count
