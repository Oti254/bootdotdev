class Wall:
    armor = 10
    height = 5

    # Double the armor
    def fortify(self):
        self.armor *= 2

"""
Methods are functions written in class declarations.
They take their first parameter as self always.
Which is the instance of the class itself.
Self is a reference of the object and can be used to update the value of properties
"""
