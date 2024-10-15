import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
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
  Input,
} from "@/components/ui/form";
import { Textarea } from "@/components/ui/textarea";
import type { CreateTaskInput } from "@/features/tasks/api/createTask";
import { useCreateTask } from "@/features/tasks/api/createTask";
import { useForm } from "react-hook-form";

export const CreateTask = () => {
  const form = useForm<CreateTaskInput>({
    defaultValues: {
      title: "",
      description: "",
    },
  });

  const createTask = useCreateTask();

  const handleSubmit = (values: CreateTaskInput) => {
    createTask.mutate(values);
  };

  return (
    <>
      <Card className="w-[400px]">
        <CardHeader className="space-y-1">
          <CardTitle className="text-2xl">タスクを作成</CardTitle>
          <CardDescription>
            タイトルと説明を入力してタスクを作成してください
          </CardDescription>
        </CardHeader>
        <CardContent className="grid gap-4">
          <Form {...form}>
            <form noValidate={true} onSubmit={form.handleSubmit(handleSubmit)}>
              <div className="grid gap-2">
                <FormField
                  control={form.control}
                  name="title"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>タイトル</FormLabel>
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
                  name="description"
                  render={({ field }) => (
                    <FormItem>
                      <FormLabel>説明</FormLabel>
                      <FormControl>
                        <Textarea {...field} />
                      </FormControl>
                      <FormMessage />
                    </FormItem>
                  )}
                />
              </div>
              <Button
                className="w-full mt-6"
                type="submit"
                disabled={createTask.isPending}
              >
                タスクを作成
              </Button>
            </form>
          </Form>
        </CardContent>
      </Card>
    </>
  );
};
