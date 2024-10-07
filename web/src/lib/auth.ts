import { api } from "@/lib/apiClient";
import type { User, UserResponse } from "@/types/api";
import axios from "axios";
import { configureAuth } from "react-query-auth";
import { z } from "zod";

const getUser = async (): Promise<User> => {
  try {
    return await api.get("/v1/users/me");
  } catch (error) {
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      return Promise.resolve({
        id: "",
        email: "",
        nickname: "",
      });
    }
    throw error;
  }
};

const logout = async (): Promise<void> => {
  return await api.post("/v1/auth/logout");
};

export const loginInputSchema = z.object({
  email: z.string().email(),
  password: z.string().min(8),
});

export type LoginInput = z.infer<typeof loginInputSchema>;
const loginWithEmailAndPassword = async (
  input: LoginInput,
): Promise<UserResponse> => {
  return await api.post("/v1/auth/login", input);
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
  return await api.post("/v1/users", input);
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
