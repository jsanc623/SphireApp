import peewee
from peewee import *
import datetime

db = MySQLDatabase('adserve', user='', passwd='', host='')

class Categories(peewee.Model):
    id = peewee.IntegerField()
    name = peewee.CharField()
    parent_id = peewee.IntegerField()
    created_date = peewee.DateTimeField(default=datetime.datetime.now)

    class Meta:
        database = db

class Loader(object):
    f_name = "cats.txt"
    f_data = None
    parent_id = None

    def __init__(self):
        pass

    def load(self):
        with open(self.f_name, "r") as fp:
            db.execute_sql("TRUNCATE categories;")
            for line in fp:
                self.parent_id = None

                # Split our line into parts
                parts = line.split("/")

                # Create our key (category name)
                key = parts[0].strip(" ")

                # Delete the category name from our parts list
                del parts[0]

                # Iterate over the remaining parts and append them to the value list
                for part in parts:
                    part = part.strip(" ").replace("\r\n", "")

                    # Set the parent to 0, unless we have a last part id
                    parent = 0
                    if self.parent_id is not None:
                        parent = self.parent_id

                    try:
                        part_found = Categories.get(name=part)
                        self.parent_id = part_found.id
                        print "PART_FOUND.ID", part_found.id, part
                    except Exception as e:
                        print "SAVE", part, parent
                        category = Categories(name=part, parent_id=parent)
                        category.save()
                        self.parent_id = category.id

L = Loader()
L.load()
