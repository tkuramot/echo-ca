import { Button } from "@/components/ui/button/Button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form/Form";
import { Input } from "@/components/ui/form/Input";
import { Link } from "@/components/ui/link";
import {
  type RegisterInput,
  registerInputSchema,
  useRegister,
} from "@/lib/auth";
import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { useSearchParams } from "react-router-dom";

type RegisterFormProps = {
  onSuccess: () => void;
};

export const RegisterForm = ({ onSuccess }: RegisterFormProps) => {
  const [searchParams] = useSearchParams();
  const redirectTo = searchParams.get("redirect-to");

  const register = useRegister({
    onSuccess,
  });
  const handleSubmit = async (values: RegisterInput) => {
    await register.mutate(values);
  };

  const form = useForm<RegisterInput>({
    resolver: zodResolver(registerInputSchema),
    defaultValues: {
      nickname: "",
      email: "",
      password: "",
    },
  });

  return (
    <>
      <Form {...form}>
        <form
          noValidate={true}
          onSubmit={form.handleSubmit(handleSubmit)}
          className="space-y-8"
        >
          <FormField
            control={form.control}
            name="nickname"
            render={({ field }) => (
              <FormItem>
                <FormLabel>ニックネーム</FormLabel>
                <FormControl>
                  <Input type="text" {...field} />
                </FormControl>
                <FormMessage />
              </FormItem>
            )}
          />
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
          <Button type="submit">登録</Button>
        </form>
      </Form>
      <div className="mt-2 flex items-center justify-end">
        <div className="text-sm">
          <Link
            to={`/auth/login${redirectTo ? `?redirectTo=${encodeURIComponent(redirectTo)}` : ""}`}
            className="font-medium text-blue-600 hover:text-blue-500"
          >
            Log In
          </Link>
        </div>
      </div>
    </>
  );
};
