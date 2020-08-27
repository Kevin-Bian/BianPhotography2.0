import React from "react";
import ReactDOM from "react-dom";
import Main from "../components/Main/Main"
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";

function App() {
  return (
    <div className="App">
      <Main />
    </div>
  );
}

export default App