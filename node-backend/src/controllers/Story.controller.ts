import { Request, Response } from 'express';

export async function AddStory(req: Request, res: Response) {
    console.log("Story Stuff")
    return res.status(200)
}
