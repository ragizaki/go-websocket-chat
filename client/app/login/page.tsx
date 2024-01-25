"use client";

import { useRouter } from "next/navigation";
import { useState, useContext, useEffect } from "react";
import { AuthContext, UserInfo } from "@/context/AuthContext";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const { authenticated } = useContext(AuthContext);
  const router = useRouter();

  useEffect(() => {
    if (authenticated) {
      router.push("/");
    }
  }, [authenticated]);

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const res = await fetch("http://localhost:8080/api/auth/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, password }),
      });
      console.log(res);
      const { username, id } = await res.json();

      if (!res.ok) {
        throw new Error("Unable to login");
      }
      const user: UserInfo = { username, id };
      localStorage.setItem("user", JSON.stringify(user));

      return router.push("/");
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="flex items-center justify-center min-w-full min-h-screen">
      <form
        className="flex flex-col w-full md:w-1/2 lg:w-1/3 space-y-3"
        onSubmit={handleLogin}
      >
        <h1 className="text-3xl font-bold text-center mb-3">
          Welcome to Chatr!
        </h1>
        <input
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          placeholder="Email"
          className="p-3 rounded-md border-2 border-gray-300 outline-none focus:border-blue-400"
        />
        <input
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          placeholder="Password"
          type="password"
          className="p-3 rounded-md border-2 border-gray-300 outline-none focus:border-blue-400"
        />
        <button
          disabled={!email || !password}
          className="bg-blue-500 p-3 rounded-md text-white enabled:hover:bg-blue-600 transition ease-in disabled:bg-blue-200"
        >
          Login
        </button>
      </form>
    </div>
  );
}
