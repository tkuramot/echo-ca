import { queryOptions, useQuery } from "@tanstack/react-query";

import type { TaskStatus } from "@/features/tasks/types";
import { api } from "@/lib/apiClient";
import type { QueryConfig } from "@/lib/reactQuery";
import type { TaskListResponse } from "@/types/api";

type TaskFilter = {
  userID?: number;
  status?: TaskStatus;
};

const getTaskList = async ({ userID, status }: TaskFilter) => {
  const response: TaskListResponse = await api.get("/v1/tasks", {
    params: {
      userID,
      status,
    },
  });
  return response.tasks;
};

export const getTaskListQueryOptions = (filter: TaskFilter = {}) => {
  return queryOptions({
    queryKey: ["taskList"],
    queryFn: () => getTaskList(filter),
  });
};

type UseTaskListOptions = {
  filter: TaskFilter;
  queryConfig?: QueryConfig<typeof getTaskListQueryOptions>;
};

export const useTaskList = ({ filter, queryConfig }: UseTaskListOptions) => {
  return useQuery({
    ...getTaskListQueryOptions(filter),
    ...queryConfig,
  });
};
