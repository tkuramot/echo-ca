import { createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/(app)/")({
  component: HomeComponent,
  beforeLoad: ({ context, location }) => {
    if (!context.user) {
      throw redirect({
        to: "/auth/login",
        search: {
          redirect: location.href,
        },
      });
    }
  },
});

function HomeComponent() {
  return (
    <div className="p-2">
      <h3>Welcome Home!</h3>
    </div>
  );
}
