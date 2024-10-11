import { type QueryClient, useQueryClient } from "@tanstack/react-query";
import { useMemo } from "react";
import { RouterProvider, createBrowserRouter } from "react-router-dom";

import { ProtectedRoute } from "@/lib/auth";

import { AppRoot } from "@/app/routes/app/AppRoot";

export const createAppRouter = (_: QueryClient) =>
  createBrowserRouter([
    {
      path: "/",
      lazy: async () => {
        const { LandingRoute } = await import("@/app/routes/LandingRoute");
        return { Component: LandingRoute };
      },
    },
    {
      path: "/auth/register",
      lazy: async () => {
        const { RegisterRoute } = await import(
          "@/app/routes/auth/RegisterRoute"
        );
        return { Component: RegisterRoute };
      },
    },
    {
      path: "/auth/login",
      lazy: async () => {
        const { LoginRoute } = await import("@/app/routes/auth/LoginRoute");
        return { Component: LoginRoute };
      },
    },
    {
      path: "/app",
      element: (
        <ProtectedRoute>
          <AppRoot />
        </ProtectedRoute>
      ),
      children: [
        {
          path: "dashboard",
          lazy: async () => {
            const { DashboardRoute } = await import(
              "@/app/routes/app/DashboardRoute"
            );
            return { Component: DashboardRoute };
          },
        },
      ],
    },
    {
      path: "*",
      lazy: async () => {
        const { NotFoundRoute } = await import("@/app/routes/NotFoundRoute");
        return { Component: NotFoundRoute };
      },
    },
  ]);

export const AppRouter = () => {
  const queryClient = useQueryClient();

  const router = useMemo(() => createAppRouter(queryClient), [queryClient]);

  return <RouterProvider router={router} />;
};
