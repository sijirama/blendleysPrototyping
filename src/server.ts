import Express  from "express"

const app = Express()

app.get("/" , (res, req) => {
    res.send("<h2>We are online bitches</h2>")
})
