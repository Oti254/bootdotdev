def destroy_walls(wall_healths):
    new_wall_healths = []
    for wall_health in wall_healths:
        if wall_health > 0:
            new_wall_healths.append(wall_health)
    return new_wall_healths 
