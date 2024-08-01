import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import "katex/dist/katex.min.css";
import Router from "./router";

ReactDOM.createRoot(document.getElementById("root")!).render(
  <React.StrictMode>
    <Router />
  </React.StrictMode>,
);
