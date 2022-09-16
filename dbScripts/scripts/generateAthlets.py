import csv
from dataclasses import dataclass, field
from tracemalloc import start

from mimesis.schema import Field, Schema
from mimesis.builtins import RussiaSpecProvider
from mimesis.enums import Gender

from random import choice, randint

def randomHand():
    data = ['right', 'left']
    return choice(data)

def randomWeaponType():
    data = ['рапира', 'сабля', 'шпага']
    return choice(data)

def randomRank():
    data = ['1', '2', '3', 'КМС', 'МС']
    return choice(data)

def competitionDescription():
    f = Field('ru')
    return {
        'id_account': randint(2419, 4817),
        'hand': randomHand(),
        'insurance': f('boolean'),
        'license': f('boolean'),
        'weapon_type': randomWeaponType(),
        'rank': randomRank()
    }
    
schema = Schema(schema=competitionDescription)

with open("./data/athlets.csv", "a", newline="") as file:
    field_names = ['id_account', 'hand', 'insurance', 'license', 'weapon_type', 'rank']
    writer = csv.DictWriter(file, fieldnames=field_names, 
                            quotechar='"', quoting=csv.QUOTE_ALL)

    # writer.writeheader()
    
    writer.writerows(schema.create(iterations=100))
