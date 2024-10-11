import { LoginForm } from "@/features/auth/components/LoginForm";
import { useNavigate, useSearchParams } from "react-router-dom";

export const LoginRoute = () => {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const redirectTo = searchParams.get("redirect-to");

  return (
    <div className="flex h-screen items-center bg-white">
      <div className="mx-auto max-w-7xl px-4 py-12 text-center sm:px-6 lg:px-8 lg:py-16">
        <LoginForm
          onSuccess={() => {
            navigate(`${redirectTo ? `${redirectTo}` : "/app"}`, {
              replace: true,
            });
          }}
        />
      </div>
    </div>
  );
};
