import { api } from "@/lib/apiClient";
import type { User, UserResponse } from "@/types/api";
import { configureAuth } from "react-query-auth";
import { z } from "zod";

export const getUser = async (): Promise<User> => {
  const response: UserResponse = await api.get("/v1/users/me");
  return response.user;
};

const logout = async (): Promise<void> => {
  return await api.post("/v1/auth/logout");
};

export const loginInputSchema = z.object({
  email: z.string().email(),
  password: z.string().min(8),
});

export type LoginInput = z.infer<typeof loginInputSchema>;
const loginWithEmailAndPassword = async (input: LoginInput): Promise<User> => {
  const response: UserResponse = await api.post("/v1/auth/login", input);
  return response.user;
};

export const registerInputSchema = z.object({
  nickname: z.string().min(2).max(255),
  email: z.string().email(),
  password: z.string().min(8),
});

export type RegisterInput = z.infer<typeof registerInputSchema>;
const registerWithEmailAndPassword = async (
  input: RegisterInput,
): Promise<User> => {
  const response: UserResponse = await api.post("/v1/users", input);
  return response.user;
};

const authConfig = {
  userFn: getUser,
  loginFn: async (data: LoginInput) => {
    return await loginWithEmailAndPassword(data);
  },
  registerFn: async (data: RegisterInput) => {
    return await registerWithEmailAndPassword(data);
  },
  logoutFn: logout,
};

export const { useUser, useLogin, useLogout, useRegister, AuthLoader } =
  configureAuth(authConfig);

export type AuthContext = {
  isAuthenticated: boolean;
  user: User | null;
};
