import Express, { Request, Response } from "express"

const app = Express()

app.get("/", (_req: Request, res: Response) => {
    res.send("<h2>We are online bitches</h2>")
})

const PORT = process.env.PORT || 3000
app.listen(PORT, () => {
    console.clear();
    console.log("Server is listening on PORT:", PORT);
})
