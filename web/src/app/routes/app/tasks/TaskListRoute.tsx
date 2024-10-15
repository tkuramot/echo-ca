import { CreateTask } from "@/features/tasks/components/CreateTask";
import { TaskList } from "@/features/tasks/components/TaskList";

export const TaskListRoute = () => {
  return (
    <div className="w-full max-w-3xl mx-auto mt-5 px-4 sm:px-6 lg:px-8">
      <div className="flex justify-center my-5">
        <CreateTask />
      </div>
      <div className="my-5">
        <TaskList />
      </div>
    </div>
  );
};
