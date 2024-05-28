import { Navbar } from "../navbar/navbar";
import "./Index.css";
export function Index() {
  return (
    <div>
      <Navbar />
      <body className="background-custom">
        <div className="container-fluid ">
          <div className="parent container d-flex align-items-center justify-content-center">
            <h1 className="title-custom">Be conscious.</h1>
          </div>
          <a
            type="button"
            className="btn btn-outline-light rounded-0 btn-custom"
            role="button"
          >
            Get started
          </a>
        </div>
      </body>
    </div>
  );
}
