import { LoginUserForm } from "@/features/auth/components/LoginUserForm";
import { createFileRoute, redirect } from "@tanstack/react-router";
import { z } from "zod";

const fallback = "/" as const;

export const Route = createFileRoute("/auth/login")({
  validateSearch: z.object({
    redirect: z.string().optional().catch(""),
  }),
  beforeLoad: ({ context, search }) => {
    if (context.auth.isAuthenticated) {
      throw redirect({
        to: search.redirect || fallback,
      });
    }
  },
  component: LoginUserPage,
});

function LoginUserPage() {
  return (
    <div className="flex justify-center items-center h-screen">
      <LoginUserForm />
    </div>
  );
}
