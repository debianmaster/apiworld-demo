const express = require('express')
const os=require('os') 
const app = express()
const port = 3000
var cors = require('cors')
app.use(cors())

app.get('/nodejs/props', (req, res) => {
  res.json({
      favColor:"lightblue",
      name:"nodejs",
      hostname: os.hostname()
  })
})

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})