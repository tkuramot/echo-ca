export type BaseEntity = {
  id: string;
};

export type Entity<T> = {
  [K in keyof T]: T[K];
} & BaseEntity;

export type User = Entity<{
  email: string;
  nickname: string;
}>;

export type UserResponse = {
  user: User;
};

export type Task = Entity<{
  title: string;
  description: string;
  status: "backlog" | "in_progress" | "done" | "canceled";
}>;

export type TaskListResponse = {
  tasks: Task[];
};

export type CreateTaskResponse = {
  task: Task;
};
