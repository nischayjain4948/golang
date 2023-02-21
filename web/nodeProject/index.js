const express = require("express");
const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.json());


app.get("/api/golang", (req,res)=>{
    console.log("api called");
    return res.status(200).json({message:"This is a get api for the route /api/golang"})
})

app.listen(PORT, ()=>{
    console.log(`Server is running on port ${PORT}`);

})