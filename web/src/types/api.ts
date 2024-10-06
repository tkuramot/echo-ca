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
