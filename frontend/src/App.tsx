import { useEffect } from "react";

import "bootstrap/dist/css/bootstrap.min.css";

function App() {
  useEffect(() => {
    fetch("http://localhost:8080/ping")
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
      });
  }, []);

  return (
    <>
      <h1> Hello World</h1>
    </>
  );
}

export default App;
