export function Placeholder(props) {
  const { id, name, type, idInput, aria, placeholder, value, func } = props;
  console.log(value);
  return (
    <div className="form-group form-group-lg">
      <label id={id}></label>
      <input
        value={value}
        onChange={func}
        name={name}
        type={type}
        className="form-control form-control-lg rounded-0 mx-auto placeholder-custom"
        id={idInput}
        aria-describedby={aria}
        placeholder={placeholder}
      />
    </div>
  );
}
