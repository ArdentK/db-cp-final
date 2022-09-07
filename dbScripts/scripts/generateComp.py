import csv
from dataclasses import dataclass, field

from mimesis.schema import Field, Schema
from mimesis.builtins import RussiaSpecProvider

from random import choice, randint

def randomAgeCategory():
    data = ['дети', 'кадеты', 'юниоры', 'u-23', 'взрослые']
    return choice(data)

def randomWeaponType():
    data = ['рапира', 'сабля', 'шпага']
    return choice(data)

def randomStatus():
    data = ['Мировые', 'Европейские', 'Российские']
    return choice(data)

def randomType():
    data = ['Олимпийские игры', 'Чемпионат мира', 'Официальные FIE', 'Гран-при',
            'Кубок мира', 'Сателлит', 'Чемпионат Европы', 'цикл U-23', 'Кадетский цикл', 
            'Чемпионат России', 'Кубок России', 'Всероссийские соревнования', 'Соревнования ФО РФ',
            'Межрегиональные соревнования', 'Соревнования субъекта РФ']
    return choice(data)

def randomSex():
    data = ['female', 'male']
    return choice(data)

def competitionDescription():
    f = Field('ru')
    return {
        'name': ' '.join(f('words', quantity=2)).capitalize(),
        'dt': f('date'),
        'age_category': randomAgeCategory(),
        'weapon_type': randomWeaponType(),
        'is_team': f('boolean'),
        'status': randomStatus(),
        'sex': randomSex(),
        'type': randomType(),
        'count': 0
    }
    
schema = Schema(schema=competitionDescription)

with open("./data/competitions.csv", "a", newline="") as file:
    field_names = ['name', 'dt', 'age_category', 'weapon_type', 'is_team', 'status', 'sex', 'type']
    writer = csv.DictWriter(file, fieldnames=field_names, 
                            quotechar='"', quoting=csv.QUOTE_ALL)

    # writer.writeheader()
    
    writer.writerows(schema.create(iterations=100))