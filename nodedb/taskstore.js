"use strict";

const mongodb = require("mongodb");

class MongoStore {
    constructor(db, colName) {
        this.collection = db.collection(colName);
    }

    insert(task) {
        task._id = new mongodb.ObjectID();
        return this.collection.insertOne(task)
    }

    async update(id, updates) {
        let updateDoc = {
            "$set": updates
        };

        return await this.collection.findOneAndUpdate({_id: id}, updateDoc, {returnOriginal: false})
    }

    async get(id) {
        return await this.collection.findOne({_id: id})
    }

    async delete(id) {
        return await this.collection.deleteOne({_id: id})
    }

    async getAll(completed) {
        return await this.collection.find({ completed })
            .limit(1000)
            .toArray()
    }
}

module.exports = MongoStore;