import random

import pymongo

client = pymongo.MongoClient("mongodb://admin:password@localhost:27017")

mydb = client["User_data"]
users = mydb.user_info

record = {"first_name" : "naga2", "age" : 31, "gender" : "male2"}
for i in range(12, 21):
    record = {"first_name": "naga"+str(i) , "age": random.randrange(1,100), "gender": random.choice(["male","female", "trans"])}
    users.insert_one(record)
