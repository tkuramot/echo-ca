import { api } from "@/lib/apiClient";
import type { User, UserResponse } from "@/types/api";
import type React from "react";
import { configureAuth } from "react-query-auth";
import { Navigate, useLocation } from "react-router-dom";

export const getUser = async (): Promise<User> => {
  const response: UserResponse = await api.get("/v1/users/me");
  return response.user;
};

const logout = async (): Promise<void> => {
  return await api.post("/v1/auth/logout");
};

export type LoginInput = {
  email: string;
  password: string;
};

const loginWithEmailAndPassword = async (input: LoginInput): Promise<User> => {
  const response: UserResponse = await api.post("/v1/auth/login", input);
  return response.user;
};

export type RegisterInput = {
  nickname: string;
  email: string;
  password: string;
};

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

export const ProtectedRoute = ({ children }: { children: React.ReactNode }) => {
  const { data: user, isFetched } = useUser();
  const location = useLocation();

  if (isFetched && !user) {
    return (
      <Navigate
        to={`/auth/login${location.pathname ? `?redirect-to=${encodeURIComponent(location.pathname)}` : ""}`}
        replace={true}
      />
    );
  }

  return children;
};
