import { Link } from "react-router-dom";

export function NavbarBtn(props) {
  const { link, name } = props;
  return (
    <Link to={link}>
      <li className="nav-item">
        <a className="nav-link">{name}</a>
      </li>
    </Link>
  );
}
