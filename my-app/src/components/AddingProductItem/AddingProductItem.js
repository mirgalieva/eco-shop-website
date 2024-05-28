import { Link } from "react-router-dom";
import "./AddingProductItem.css";

export function AddingProductItem(props) {
  const { imageUrl, product, description, price, id } = props;
  return (
    <div>
      <div className="col">
        <Link to={"/product/" + id}>
          <div className="card rounded-0 card-custom">
            <img className="card-img-top image-custom" src={imageUrl} alt="" />
            <div className="card-body">
              <h5 className="card-title">{product}</h5>
              <p className="card-text">{description}</p>
              <a className="btn btn-primary btn-price">${price}</a>
            </div>
          </div>
        </Link>
      </div>
    </div>
  );
}
