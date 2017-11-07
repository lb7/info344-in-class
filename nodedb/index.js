"use strict";

const express = require("express");
const app = express();

const mongodb = require('mongodb');
const MongoStore = require('./taskstore');

const addr = process.env.ADDR || "localhost:4000";
const [host, port] = addr.split(":");

const mongoAddr = process.env.DBADDR || "localhost:27017";
const mongoURL = `mongodb://${mongoAddr}/tasks`;

mongodb.MongoClient.connect(mongoURL).then(db => {
    let store = new MongoStore(db, 'tasks');

    app.use(express.json());

    app.post('/v1/tasks', async (req, res) => {
        //insert new task
        try {
            let task = {
                title: req.body.title,
                completed: false
            };
            let result = await store.insert(task);
            res.json(result.ops[0])
        } catch (e) {
            throw e;
        }
    });

    app.get('/v1/tasks', (req, res) => {
        //return all non-completed tasks
    });

    app.patch('/v1/tasks/:taskId', (req, res) => {
        let taskIdToFetch = req.params.taskId;
    });

    app.listen(port, host, () => {
        console.log(`server is listening at http://${addr}...`)
    });
}).catch(console.error);
