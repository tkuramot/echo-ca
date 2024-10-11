import { AppProvider } from "@/app/AppProvider";
import { AppRouter } from "@/app/AppRouter";

export const App = () => {
  return (
    <AppProvider>
      <AppRouter />
    </AppProvider>
  );
};
