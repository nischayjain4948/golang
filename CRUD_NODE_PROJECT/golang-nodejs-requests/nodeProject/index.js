const express = require("express");
const app = express();
const PORT = process.env.PORT || 3000;

app.use(express.json());


app.get("/api/golang", (req,res)=>{
    console.log("api called");
    return res.status(200).json({message:"This is a get api for the route /api/golang"})
})


app.post("/api/golang/post", (req,res)=>{
    console.log("post api called...");
    res.status(200).json({message: req.body});
})


app.post("/api/golang/formdata", (req,res)=>{
    console.log("form Data api called..");

    res.status(200).json({message : req.body})
})


app.get("/api/golang/json", async (req,res)=>{
   const jsonResponse = await fetch("https://jsonplaceholder.typicode.com/posts")
   const data = await jsonResponse.json();
   console.log("api called...");
   res.status(200).json({data});

})



app.listen(PORT, ()=>{
    console.log(`Server is running on port ${PORT}`);

})


