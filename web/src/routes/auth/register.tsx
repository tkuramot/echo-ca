import { RegisterUserForm } from "@/features/auth/components/RegisterUserForm";
import { createFileRoute } from "@tanstack/react-router";

export const Route = createFileRoute("/auth/register")({
  component: RegisterUserPage,
});

function RegisterUserPage() {
  return (
    <div className="flex justify-center items-center h-screen">
      <RegisterUserForm />
    </div>
  );
}
