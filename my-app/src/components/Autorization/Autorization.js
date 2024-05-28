import "./Autorization.css";
import { useState } from "react";
import { Form } from "../Form/Form";

export function Autorization() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleChangeEmail = (event) => {
    setEmail(event.target.value);
  };

  const handleChangePassword = (event) => {
    setPassword(event.target.value);
  };

  const placeholderItems = [
    {
      id: "email",
      value: email,
      func: handleChangeEmail,
      name: "email",
      type: "email",
      placeholder: "Enter you email...",
      idInput: "exampleInputEmail1",
      aria: "emailHelp",
    },

    {
      id: "password",
      value: password,
      func: handleChangePassword,
      name: "password",
      type: "password",
      placeholder: "Enter password...",
      idInput: "exampleInputPassword1",
      aria: "passwordHelp",
    },
  ];

  const handleAdd = async (event) => {
    event.preventDefault();
    const values = {
      email,
      password,
    };
    await fetch(/*"http://localhost:3001/users"*/"http://localhost:8080/users/register", {
      method: "POST",
      body: JSON.stringify(values),
      headers: {
        "Content-Type": "application/json",
      },
    });
  };

  return (
    <Form
      title="Sign in"
      placeholderItems={placeholderItems}
      btnName="Continue"
      handleAdd={handleAdd}
    />
  );
}
