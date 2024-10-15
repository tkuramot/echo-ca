import { getTaskListQueryOptions } from "@/features/tasks/api/getTaskList";
import { api } from "@/lib/apiClient";
import type { MutationConfig } from "@/lib/reactQuery";
import type { CreateTaskResponse } from "@/types/api";
import { useMutation, useQueryClient } from "@tanstack/react-query";

export type CreateTaskInput = {
  title: string;
  description: string;
};

const createTask = async (input: CreateTaskInput) => {
  const response: CreateTaskResponse = await api.post("/v1/tasks", input);
  return response.task;
};

type UseCreateTaskOptions = {
  mutationConfig?: MutationConfig<typeof createTask>;
};

export const useCreateTask = ({
  mutationConfig,
}: UseCreateTaskOptions = {}) => {
  const queryClient = useQueryClient();
  const { onSuccess, ...restConfig } = mutationConfig ?? {};

  return useMutation({
    onSuccess: (...args) => {
      const _ = queryClient.invalidateQueries({
        queryKey: getTaskListQueryOptions().queryKey,
      });
      onSuccess?.(...args);
    },
    ...restConfig,
    mutationFn: createTask,
  });
};
