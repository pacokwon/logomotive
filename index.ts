import express from 'express';

const app = express();

app.get('/', (req, res) => {
  res.json({});
});

app.listen(8000, () => {
  console.log('App listening on port 8000');
});
