const express = require("express");
const app = express();
const port = 3000;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).send("Welcome to my first express app");
});

app.get('/get', (req, res) => {
  res.status(200).json({ message: "Hello from learnCodeonline.in" });
});

app.post('/post', (req, res) => {
  let myJson = req.body;
  res.status(200).json(myJson); // Using .json to automatically stringify the object
});

app.post('/postform', (req, res) => {
  let myJson = req.body;
  res.status(200).send(JSON.stringify(myJson)); // Using .send to send the object as a string
});

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`);
});
