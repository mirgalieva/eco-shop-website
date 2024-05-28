import "./Form.css";
import { Navbar } from "../navbar/navbar";
import { Placeholder } from "../Placeholder/Placeholder";

export function Form(props) {
  const { title, placeholderItems, btnName, handleAdd } = props;
  return (
    <div>
      <body background="images/david-clode-xNSCi_K179c-unsplash.jpg">
        <Navbar />
        <div className="container text-center my-5">
          <div className="row">
            <div className="col-lg-6 mx-auto container-custom">
              <div className="opacity-50">
                <form className="card p-3 bg-light rounded-0">
                  <h1>{title}</h1>
                  {placeholderItems.map((item, index) => (
                    <Placeholder
                      key={index}
                      id={item.id}
                      value={item.value}
                      func={item.func}
                      name={item.name}
                      type={item.type}
                      placeholder={item.placeholder}
                      idInput={item.idInput}
                      aria={item.aria}
                    />
                  ))}
                  <button
                    onClick={handleAdd}
                    type="submit"
                    href="/"
                    id="add-button"
                    className="add-btn btn btn-outline-light  rounded-0   mx-auto opacity-100"
                  >
                    <div>{btnName}</div>
                  </button>
                </form>
              </div>
            </div>
          </div>
        </div>
      </body>
    </div>
  );
}
