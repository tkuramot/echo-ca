import { RegisterForm } from "@/features/auth/components/RegisterForm";
import { useNavigate, useSearchParams } from "react-router-dom";

export const RegisterRoute = () => {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const redirectTo = searchParams.get("redirect-to");

  return (
    <div className="flex h-screen items-center bg-white">
      <div className="mx-auto max-w-7xl px-4 py-12 text-center sm:px-6 lg:px-8 lg:py-16">
        <RegisterForm
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
