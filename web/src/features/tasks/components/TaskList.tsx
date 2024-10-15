import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useTaskList } from "@/features/tasks/api/getTaskList";
import { useUser } from "@/lib/auth";

export const TaskList = () => {
  const { data: user } = useUser();
  const { data: tasks } = useTaskList({
    filter: {
      userID: user?.id,
    },
    queryConfig: {
      enabled: !!user,
    },
  });

  return (
    <Table>
      <TableCaption>Task List</TableCaption>
      <TableHeader>
        <TableRow>
          <TableHead className="w-[100px]">タイトル</TableHead>
          <TableHead>説明</TableHead>
          <TableHead>ステータス</TableHead>
        </TableRow>
      </TableHeader>
      <TableBody>
        {tasks?.map((task) => (
          <TableRow key={task.id}>
            <TableCell>{task.title}</TableCell>
            <TableCell>{task.description}</TableCell>
            <TableCell>{task.status}</TableCell>
          </TableRow>
        ))}
      </TableBody>
    </Table>
  );
};
