import csv
from dataclasses import dataclass, field
from tracemalloc import start

from mimesis.schema import Field, Schema
from mimesis.builtins import RussiaSpecProvider
from mimesis.enums import Gender

from random import choice, randint

def randomSex():
    data = ['female', 'male']
    return choice(data)


def competitionDescription():
    f = Field('ru')
    return {
        'name': f('full_name', gender=Gender.MALE),
        'birthday': f('date', start=2000, end=2009),
        'sex': 'male',
        'email': f('email')
    }
    
schema = Schema(schema=competitionDescription)

with open("./data/accounts.csv", "a", newline="") as file:
    field_names = ['name', 'birthday', 'sex', 'email']
    writer = csv.DictWriter(file, fieldnames=field_names, 
                            quotechar='"', quoting=csv.QUOTE_ALL)

    # writer.writeheader()
    
    writer.writerows(schema.create(iterations=100))
