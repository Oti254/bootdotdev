class Wizard:
    def __init__(self, name, stamina, intelligence):
        self.name = name
        self.__stamina = stamina
        self.__intelligence = intelligence
        self.health = self.__stamina * 100 
        self.mana = self.__intelligence * 10

"""
Encapsulation is practice of hiding the functionality in a 'black box'
We use private data members denoted by __ to encapsulate logic and data
"""
