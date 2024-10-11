import ReactDOM from "react-dom/client";
import "@/index.css";
import { App } from "@/app";
import React from "react";

const rootElement = document.getElementById("app");

if (rootElement && !rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <React.StrictMode>
      <App />
    </React.StrictMode>,
  );
}
