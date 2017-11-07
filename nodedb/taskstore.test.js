"use strict";

const MongoStore = require('./taskstore');
const mongodb = require('mongodb');

const mongoAddr = process.env.DBADDR || "localhost:27017";
const mongoURL = `mongodb://${mongoAddr}/tasks`;

describe('Mongo Task Store', () => {
    test("CRUD Cycle", async () => {
        let db = await mongodb.MongoClient.connect(mongoURL);
        let store = new MongoStore(db, 'tasks');
        let task = {
            title: "Learn Node.js to mongodb",
            tags: ["mongodb", "nodejs"]
        };

        try {
            let result = await store.insert(task);
            expect(result.insertedId).toBeDefined();

            let taskId = result.insertedId;

            result = await store.get(taskId);
            expect(result).toEqual(task);

            await store.delete(taskId);

            result = await store.get(taskId);
            expect(result).toBeFalsy()
        } catch (err) {
            throw err;
        }
    });

});