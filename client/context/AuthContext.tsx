"use client";

import { createContext, useState, useEffect } from "react";
import { useRouter } from "next/navigation";

export interface UserInfo {
  username: string;
  id: string;
}

export const AuthContext = createContext<{
  authenticated: boolean;
  setAuthenticated: React.Dispatch<React.SetStateAction<boolean>>;
  user: UserInfo;
  setUser: React.Dispatch<React.SetStateAction<UserInfo>>;
}>({
  authenticated: false,
  setAuthenticated: () => {},
  user: { username: "", id: "" },
  setUser: () => {},
});

export default function AuthContextProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [authenticated, setAuthenticated] = useState(false);
  const [user, setUser] = useState<UserInfo>({ username: "", id: "" });
  const router = useRouter();

  useEffect(() => {
    const user = localStorage.getItem("user");
    if (!user) {
      if (window.location.href !== "/signup") {
        router.push("/login");
      }
      return;
    }
    const userInfo: UserInfo = JSON.parse(user);
    if (user) {
      setAuthenticated(true);
      setUser({ username: userInfo.username, id: userInfo.id });
    }
  }, [authenticated]);

  return (
    <AuthContext.Provider
      value={{
        authenticated,
        setAuthenticated,
        user,
        setUser,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
