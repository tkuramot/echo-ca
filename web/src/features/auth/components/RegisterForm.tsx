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
} from "@/components/ui/form/Form";
import { Input } from "@/components/ui/form/Input";
import { type RegisterInput, useRegister } from "@/lib/auth";
import { useForm } from "react-hook-form";
import { useNavigate, useSearchParams } from "react-router-dom";

type RegisterFormProps = {
  onSuccess: () => void;
};

export const RegisterForm = ({ onSuccess }: RegisterFormProps) => {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const redirectTo = searchParams.get("redirect-to");

  const register = useRegister({
    onSuccess,
  });
  const handleSubmit = async (values: RegisterInput) => {
    await register.mutate(values);
  };

  const form = useForm<RegisterInput>({
    defaultValues: {
      nickname: "",
      email: "",
      password: "",
    },
  });

  return (
    <>
      <Card className="w-[400px]">
        <CardHeader className="space-y-1">
          <CardTitle className="text-2xl">アカウント登録</CardTitle>
          <CardDescription>
            必要な情報を入力して、新しいアカウントを作成してください
          </CardDescription>
        </CardHeader>
        <CardContent className="grid gap-4">
          <Form {...form}>
            <form noValidate={true} onSubmit={form.handleSubmit(handleSubmit)}>
              <div className="grid gap-2">
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
              </div>
              <div className="grid gap-2 mt-4">
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
                    </FormItem>
                  )}
                />
              </div>
              <Button
                className="w-full mt-6"
                type="submit"
                disabled={register.isPending}
              >
                登録
              </Button>
            </form>
          </Form>
        </CardContent>
        <CardFooter>
          <Button
            onClick={() => {
              navigate(
                `/auth/login${redirectTo ? `?redirect-to=${redirectTo}` : ""}`,
              );
            }}
            variant="link"
            className="px-0 text-sm text-muted-foreground"
          >
            すでにアカウントをお持ちですか？ ログイン
          </Button>
        </CardFooter>
      </Card>
    </>
  );
};
