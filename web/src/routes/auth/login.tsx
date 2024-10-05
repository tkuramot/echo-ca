import { LoginUserPage } from "@/features/auth/login";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/auth/login")({
  component: LoginUserPage,
});
