class Character(object):
    def __init__(self):
        self.icon = ''
        self.name = ''
        self.number = 0

    def add_number(self):
        self.number += 1

    def reduce_number(self):
        self.number -= 1

    def set_number(self, num):
        self.number = num

class Werewolves(Character):
    def __init__(self):
        self.icon = 'werewolves'
        self.name = 'werewolves'

class Seer(Character):
    def __init__(self):
        self.icon = 'seer'
        self.name = 'seer'
