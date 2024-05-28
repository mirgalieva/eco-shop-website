import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { AddingProductItem } from "../AddingProductItem/AddingProductItem";

export function Product() {
  const params = useParams();
  const [publication, setPublication] = useState(null);
  useEffect(() => {
    (async () => {
      const data = await fetch(`http://localhost:3001/post/${params.id}`);
      const post = await data.json();
      setPublication(post);
    })();
  }, [params.id]);
  if (!publication) {
    return <div>Downloading</div>;
  }

  return (
    <AddingProductItem
      id={publication.id}
      product={publication.product}
      price={publication.price}
      description={publication.description}
      imageUrl={publication.imageUrl}
    />
  );
}
