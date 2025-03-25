const express = require("express");
const axios = require("axios");

const app = express();
const PORT = 8080;

app.get("/score", async (req, res) => {
    const node = req.query.node;
    
    let energyUsage = Math.floor(Math.random() * 100); 
    let score = 100 - energyUsage;

    console.log(`Node: ${node}, Energy Usage: ${energyUsage}, Score: ${score}`);

    res.json({ node, score });
});

// Start the server
app.listen(PORT, () => {
    console.log(`Middleware is running on http://localhost:${PORT}`);
});
