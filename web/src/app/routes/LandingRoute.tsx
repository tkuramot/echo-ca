import { useNavigate } from "react-router-dom";

import { Button } from "@/components/ui/button";
import { useUser } from "@/lib/auth";

export const LandingRoute = () => {
  const navigate = useNavigate();
  const { data: user } = useUser();

  const handleStart = () => {
    if (user) {
      navigate("/app/dashboard");
    } else {
      navigate("/auth/login");
    }
  };

  return (
    <>
      <div className="flex h-screen items-center bg-white">
        <div className="mx-auto max-w-7xl px-4 py-12 text-center sm:px-6 lg:px-8 lg:py-16">
          <Button onClick={handleStart}>Get started</Button>
        </div>
      </div>
    </>
  );
};
