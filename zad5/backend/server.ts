import express from "express";
import cors from "cors";
import { products } from "./products.ts";

const app = express();

app.use(cors());
app.use(express.json());

// GET products
app.get("/products", (req, res) => {
  res.json(products);
});

// POST payment
app.post("/payments", (req, res) => {
  console.log("Payment:", req.body);
  res.json({ status: "ok" });
});

app.listen(8080, () => {
  console.log("Backend running on port 8080");
});