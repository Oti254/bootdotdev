class Unit:
    def __init__(self, name, pos_x, pos_y):
        self.name = name
        self.pos_x = pos_x
        self.pos_y = pos_y

    def in_area(self, x_1, y_1, x_2, y_2):
        min_x, max_x = min(x_1, x_2), max(x_1, x_2)
        min_y, max_y = min(y_1, y_2), max(y_1, y_2)
        return (min_x <= self.pos_x <= max_x) and (min_y <= self.pos_y <= max_y)

class Dragon(Unit):
    def __init__(self, name, pos_x, pos_y, fire_range):
        super().__init__(name, pos_x, pos_y)
        self.__fire_range = fire_range

    def breathe_fire(self, x, y, units):
        blast_hit = []
        x1 = x + self.__fire_range
        x2 = x - self.__fire_range
        y1 = y + self.__fire_range
        y2 = y - self.__fire_range
        for unit in units:
            if unit.in_area(x1, y1, x2, y2):
                blast_hit.append(unit)
        return blast_hit
