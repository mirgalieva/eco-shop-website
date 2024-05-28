import "./App.css";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import React from "react";
import { Main } from "./components/Main/Main";
import { Index } from "./components/Index/Index";
import { Add } from "./components/AddingForm/AddingForm.js";
import { Autorization } from "./components/Autorization/Autorization";
import { Product } from "./components/Product/Product";
function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Main />} />
        <Route path="/about" element={<Index />} />
        <Route path="/add" element={<Add />} />
        <Route path="/autorization" element={<Autorization />} />
        <Route path="/product/:id" element={<Product />} />
      </Routes>
    </BrowserRouter>
  );
}
export default App;
