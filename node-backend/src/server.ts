import express from 'express';
import { CONFIG } from './lib/config';
import { InitializeServer } from './init';
import { DatabaseConnect } from './lib/database';

let server = express();
console.log('runing');

//NOTE: try and connect to the database
DatabaseConnect();

//NOTE: initialize the server pls
server = InitializeServer(server);

server.listen(CONFIG.PORT, () => {
    console.log('listening on port ' + CONFIG.PORT);
});
