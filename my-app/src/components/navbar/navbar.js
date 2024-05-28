import { Link } from "react-router-dom";
import { NavbarBtn } from "../NavbarBtn/NavbarBtn";

export function Navbar() {
  const btnItems = [
    {
      link: "/add",
      name: "Add new product",
    },

    {
      link: "/about",
      name: "About us",
    },

    {
      link: "/",
      name: "Search",
    },
    {
      link: "/autorization",
      name: "Get started",
    },
  ];

  return (
    <div>
      <link
        href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css"
        rel="stylesheet"
        integrity="sha384-T3c6CoIi6uLrA9TneNEoa7RxnatzjcDSCmG1MXxSR1GAsXEV/Dwwykc2MPK8M2HN"
        crossorigin="anonymous"
      />
      <nav className="navbar navbar-expand-lg navbar-dark background-color-custom">
        <div className="container-fluid">
          <Link to="/about">
            <a className="navbar-brand">ECOSHOP</a>
            <button
              className="navbar-toggler"
              type="button"
              data-bs-toggle="collapse"
              data-bs-target="#navbarNav"
              aria-controls="navbarNav"
              aria-expanded="false"
              aria-label="Toggle navigation"
            >
              <span className="navbar-toggler-icon"></span>
            </button>
          </Link>

          <div className="text-end">
            <div className="collapse navbar-collapse" id="navbarNav">
              <ul className="navbar-nav">
                {btnItems.map((item, index) => (
                  <NavbarBtn key={index} link={item.link} name={item.name} />
                ))}
              </ul>
            </div>
          </div>
        </div>
      </nav>
    </div>
  );
}
