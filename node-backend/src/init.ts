import express, { Express } from 'express';
import StoryRouter from './routes/story.route';
// import {  } from "cors" d.t.s issue

export function InitializeServer(server: Express): Express {
    //NOTE: health check, for logging purposes, to see if the server is still on or something
    server.get('/healthcheck', (_req, res) => {
        res.sendStatus(200);
    });

    // NOTE: middleware comes here
    server.use(express.json());
    server.use(express.urlencoded({ extended: false }));

    // NOTE: routes comes here
    server.use('/', StoryRouter);

    // NOTE: catch all exception
    server.all('*', (req, res) => {
        res.status(404).json({ error: `Route ${req.originalUrl} not found` });
    });

    return server;
}
