#!/bin/sh

mongoimport --host mongo --port 27017 --db menu --collection deserts --mode upsert --type json --file /seeds/deserts.json --jsonArray
mongoimport --host mongo --port 27017 --db menu --collection drinks --mode upsert --type json --file /seeds/drinks.json --jsonArray
mongoimport --host mongo --port 27017 --db menu --collection foods --mode upsert --type json --file /seeds/foods.json --jsonArray