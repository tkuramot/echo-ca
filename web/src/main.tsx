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
    // biome-ignore lint/style/noNonNullAssertion: this is guaranteed to be set by the AppProvider
    auth: undefined!,
  },
});

// Register things for typesafety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const App = () => {
  const { data: user, isFetched } = useUser();
  return (
    isFetched && <RouterProvider
      router={router}
      context={{
        auth: {
          isAuthenticated: !!user,
          user,
        },
      }}
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
