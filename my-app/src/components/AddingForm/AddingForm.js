import "./AddingForm.css";
import { useState } from "react";
import { Form } from "../Form/Form";

export function Add() {
  const [name, setName] = useState("");
  const [Price, setPrice] = useState("");
  const [Description, setDescription] = useState("");

  const handleChangeProduct = (event) => {
    setName(event.target.value);
  };

  const handleChangePrice = (event) => {
    setPrice(event.target.value);
  };

  const handleChangeDescription = (event) => {
    setDescription(event.target.value);
  };

  const placeholderItems = [
    {
      id: "name",
      value: name,
      func: handleChangeProduct,
      name: "name",
      type: "name",
      placeholder: "Enter product name...",
      idInput: "exampleInputProduct1",
      aria: "nameHelp",
    },

    {
      id: "Price",
      value: Price,
      func: handleChangePrice,
      name: "Price",
      type: "Price",
      placeholder: "Enter price...",
      idInput: "exampleInputPrice1",
      aria: "PriceHelp",
    },

    {
      id: "Description",
      value: Description,
      func: handleChangeDescription,
      name: "Description",
      type: "Description",
      placeholder: "Enter description...",
      idInput: "exampleInputDescription1",
      aria: "DescriptionHelp",
    },
  ];

  const handleAdd = async (event) => {
    event.preventDefault();
    const values = {
      Price,
      name,
      Description,
    };
    await fetch(/*"http://localhost:3001/posts"*/"http://localhost:8080/products/add", {
      method: "POST",
      body: JSON.stringify(values),
      headers: {
        "Content-Type": "application/json",
      },
    });
  };

  return (
    <Form
      title="Add new product"
      placeholderItems={placeholderItems}
      btnName="Add product"
      handleAdd={handleAdd}
    />
  );
}
