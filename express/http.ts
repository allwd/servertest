import express from "express";

const app = express();
const port = 3000;

app.get("/test", (req, res) => {
  if (req.headers.authorization !== "Test") {
    res.send(JSON.stringify({ err: "Unauthorized." }));
    return;
  }
  res.send(JSON.stringify({ status: "Success." }));
});

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`);
});
