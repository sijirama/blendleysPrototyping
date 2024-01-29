import { Router } from 'express';
import * as StoryController from '../controllers/Story.controller';

const route = Router();

route.get('/story', StoryController.AddStory);
export default route;
