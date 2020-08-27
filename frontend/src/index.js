import React, { useState, useCallback } from "react";

import ReactDOM from 'react-dom';
import { BrowserRouter } from 'react-router-dom';
import App from "./app/App"

ReactDOM.render((
  <BrowserRouter>
    <App /> {/* The various pages will be displayed by the `Main` component. */}
  </BrowserRouter>
  ), document.getElementById('root')
);