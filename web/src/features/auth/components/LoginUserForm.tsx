import { Button } from "@/components/ui/button/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form/form";
import { Input } from "@/components/ui/form/input";
import { type LoginInput, loginInputSchema, useLogin } from "@/lib/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { useNavigate, useRouter, useSearch } from "@tanstack/react-router";
import { useForm } from "react-hook-form";

const fallback = "/" as const;

export const LoginUserForm = () => {
  const router = useRouter();
  const navigate = useNavigate();
  const search = useSearch({
    from: "/auth/login",
  });

  const form = useForm<LoginInput>({
    resolver: zodResolver(loginInputSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const login = useLogin();
  const handleSubmit = async (values: LoginInput) => {
    await login.mutate(values);
    await navigate({ to: search.redirect || fallback });
  };

  return (
    <Form {...form}>
      <form
        noValidate={true}
        onSubmit={form.handleSubmit(handleSubmit)}
        className="space-y-8"
      >
        <FormField
          control={form.control}
          name="email"
          render={({ field }) => (
            <FormItem>
              <FormLabel>メールアドレス</FormLabel>
              <FormControl>
                <Input type="email" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel>パスワード</FormLabel>
              <FormControl>
                <Input type="password" {...field} />
              </FormControl>
              <FormMessage />
            </FormItem>
          )}
        />
        <Button type="submit">ログイン</Button>
      </form>
    </Form>
  );
};
