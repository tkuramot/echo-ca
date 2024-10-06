import { api } from "@/lib/apiClient";
import type { User, UserResponse } from "@/types/api";
import { configureAuth } from "react-query-auth";
import { z } from "zod";

const getUser = async (): Promise<User> => {
  const response = await api.get("/users/me");
  return response.data;
};

const logout = async () => {
  await api.post("/auth/logout");
};

export const loginInputSchema = z.object({
  email: z.string().email(),
  password: z.string().min(8),
});

export type LoginInput = z.infer<typeof loginInputSchema>;
const loginWithEmailAndPassword = async (
  input: LoginInput,
): Promise<UserResponse> => {
  return await api.post("/auth/login", input);
};

export const registerInputSchema = z.object({
  nickname: z.string().min(2).max(255),
  email: z.string().email(),
  password: z.string().min(8),
});

export type RegisterInput = z.infer<typeof registerInputSchema>;
const registerWithEmailAndPassword = async (
  input: RegisterInput,
): Promise<UserResponse> => {
  return await api.post("/users", input);
};

const authConfig = {
  userFn: getUser,
  loginFn: async (data: LoginInput) => {
    const response = await loginWithEmailAndPassword(data);
    return response.user;
  },
  registerFn: async (data: RegisterInput) => {
    const response = await registerWithEmailAndPassword(data);
    return response.user;
  },
  logoutFn: logout,
};

export const { useUser, useLogin, useLogout, useRegister, AuthLoader } =
  configureAuth(authConfig);
