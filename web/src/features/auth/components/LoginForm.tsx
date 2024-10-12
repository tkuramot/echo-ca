import { Button } from "@/components/ui/button/Button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/form/Input";
import { Link } from "@/components/ui/link";
import { type LoginInput, useLogin } from "@/lib/auth";
import { useForm } from "react-hook-form";
import { useNavigate, useSearchParams } from "react-router-dom";

type LoginFormProps = {
  onSuccess: () => void;
};

export const LoginForm = ({ onSuccess }: LoginFormProps) => {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const redirectTo = searchParams.get("redirect-to");

  const login = useLogin({
    onSuccess,
  });
  const handleSubmit = async (values: LoginInput) => {
    await login.mutate(values);
  };

  const form = useForm<LoginInput>({
    defaultValues: {
      email: "",
      password: "",
    },
  });

  return (
    <>
      <Card className="w-[400px]">
        <CardHeader className="space-y-1">
          <CardTitle className="text-2xl">ログイン</CardTitle>
          <CardDescription>
            メールアドレスとパスワードを入力してログインしてください
          </CardDescription>
        </CardHeader>
        <CardContent className="grid gap-4">
          <Form {...form}>
            <form noValidate={true} onSubmit={form.handleSubmit(handleSubmit)}>
              <div className="grid gap-2">
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
              </div>
              <div className="grid gap-2 mt-4">
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
                      <Link className="text-sm text-right" to="#">
                        パスワードを忘れましたか？
                      </Link>
                    </FormItem>
                  )}
                />
              </div>
              <Button
                className="w-full mt-6"
                type="submit"
                disabled={login.isPending}
              >
                ログイン
              </Button>
            </form>
          </Form>
        </CardContent>
        <CardFooter>
          <Button
            onClick={() => {
              navigate(
                `/auth/register${redirectTo ? `?redirectTo=${encodeURIComponent(redirectTo)}` : ""}`,
              );
            }}
            variant="link"
            className="px-0 text-sm text-muted-foreground"
          >
            アカウントを作成する
          </Button>
        </CardFooter>
      </Card>
    </>
  );
};
