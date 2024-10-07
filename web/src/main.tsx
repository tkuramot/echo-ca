import { RouterProvider, createRouter } from "@tanstack/react-router";
import ReactDOM from "react-dom/client";
import { routeTree } from "./routeTree.gen";
import "@/index.css";
import { useUser } from "@/lib/auth";
import { AppProvider } from "@/provider";
import React from "react";

// Set up a Router instance
const router = createRouter({
  routeTree,
  defaultPreload: "intent",
  context: {
    loaded: false,
    // biome-ignore lint/style/noNonNullAssertion: this is guaranteed to be set by the AppProvider
    user: undefined!,
  },
});

// Register things for typesafety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const App = () => {
  const { data, isFetched } = useUser();
  return (
    <RouterProvider
      router={router}
      context={{ loaded: isFetched, user: data?.user }}
    />
  );
};

const rootElement = document.getElementById("app");

if (rootElement && !rootElement.innerHTML) {
  const root = ReactDOM.createRoot(rootElement);
  root.render(
    <React.StrictMode>
      <AppProvider>
        <App />
      </AppProvider>
    </React.StrictMode>,
  );
}
