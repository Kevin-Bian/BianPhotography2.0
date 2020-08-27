import React from "../../node_modules/react";
import ReactDOM from "../../node_modules/react-dom";
import Main from "../components/Main/Main"
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "../../node_modules/react-router-dom";

function App() {
  return (
    <div className="App">
      <Main />
    </div>
  );
}

export default App