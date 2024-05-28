import "./Main.css";
import { Navbar } from "../navbar/navbar";
import { AddingProductItem } from "../AddingProductItem/AddingProductItem";
import { useEffect, useState } from "react";
const publicationItems = [
  {
    imageUrl: "images/ecopanda-oXbU5Z3p3r8-unsplash.jpg",
    Name: "Reusable cotton pads",
    Description:
      "co-Friendly Care: Foster a sustainable lifestyle with biodegradable bamboo toothbrushes and reusable cotton rounds.",
    Price: 200,
    id: 1,
  },
  {
    imageUrl: "images/sara-groblechner-7TgbRVEYdYY-unsplash.jpg",
    Name: "Wooden toothbrushes",
    Description:
      "Charcoal wooden toothbrushes-- no , the environmental toothbrush can be disposed safely by returning it to earth in compost or landfill.",
    Price: 200,
    id: 2,
  },
  {
    imageUrl: "images/sara-groblechner-h10-NImYZHs-unsplash.jpg",
    Name: "Thermos",
    Description:
      "Thermos Vacuum Insulated Mobile Mug JNL Series Amazon.co.jp Limited Color Model",
    Price: 200,
    id: 3,
  },
  {
    imageUrl: "images/sara-groblechner-Q0eia3UG5TQ-unsplash.jpg",
    Name: "Reusable straw",
    Description:
      "Thermos Vacuum Insulated Mobile Mug JNL Series Amazon.co.jp Limited Color Model",
    Price: 200,
    id: 4,
  },
];
export function Main() {
  const [publications, setPublications] = useState(publicationItems);
  useEffect(() => {
    (async () => {
      const data = await fetch("http://localhost:8080/products");
      const posts = await data.json();
      setPublications((oldPublications) => [
        ...oldPublications,
        ...posts.map((post) => ({
          ...post,
          imageUrl: "images/ecopanda-oXbU5Z3p3r8-unsplash.jpg",
        })),
      ]);
    })();
  }, []);

  return (
    <div>
      <div className="background-color-custom">
        <Navbar />
        <div className="container mx-auto">
          <div className="row row-cols-1 row-cols-md-2 row-cols-lg-2 g-5 mx-auto">
            {publications.map((item, index) => (
              <AddingProductItem
                key={index}
                imageUrl={item.imageUrl}
                description={item.Description}
                price={item.Price}
                product={item.Name}
                id={item.id}
              />
            ))}
          </div>
        </div>
        <main className="main"></main>
      </div>
    </div>
  );
}
