"use client";

import { useRouter } from "next/navigation";
import Link from "next/link";
import { useState, useContext, useEffect } from "react";
import { AuthContext, UserInfo } from "@/context/AuthContext";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const { authenticated } = useContext(AuthContext);
  const router = useRouter();

  // useEffect(() => {
  //   if (authenticated) {
  //     router.push("/");
  //   }
  // }, [authenticated]);

  const handleSignup = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setEmail("");
    setUsername("");
    setPassword("");

    try {
      const res = await fetch("http://localhost:8080/api/auth/signup", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ email, username, password }),
      });
      console.log(res);

      const { username: usernameRes, id } = await res.json();

      if (!res.ok) {
        throw new Error("Unable to login");
      }
      const user: UserInfo = { username: usernameRes, id };
      localStorage.setItem("user", JSON.stringify(user));

      return router.push("/");
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="flex min-h-full flex-1 flex-col justify-center items-center px-6 py-12 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-sm">
        <img
          className="mx-auto h-10 w-auto"
          src="../../favicon.ico"
          alt="Your Company"
        />
        <h2 className="mt-10 text-center text-2xl font-bold leading-9 tracking-tight text-gray-900">
          Create an account
        </h2>
      </div>

      <div className="mt-10 w-full sm:max-w-sm">
        <form className="space-y-6" onSubmit={handleSignup}>
          <div>
            <label
              htmlFor="email"
              className="block text-sm font-medium leading-6 text-gray-900"
            >
              Email address
            </label>
            <div className="mt-2">
              <input
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                type="email"
                required
                className="block w-full rounded-md ring-1 ring-gray-300 py-2 px-3 text-gray-900 shadow-sm focus:ring-blue-600 focus:outline-none"
              />
            </div>
          </div>

          <div>
            <label
              htmlFor="username"
              className="block text-sm font-medium leading-6 text-gray-900"
            >
              Username
            </label>
            <div className="mt-2">
              <input
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                required
                className="block w-full rounded-md ring-1 ring-gray-300 py-2 px-3 text-gray-900 shadow-sm focus:ring-blue-600 focus:outline-none"
              />
            </div>
          </div>

          <div>
            <div className="flex items-center justify-between">
              <label
                htmlFor="password"
                className="block text-sm font-medium leading-6 text-gray-900"
              >
                Password
              </label>
            </div>
            <div className="mt-2">
              <input
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                type="password"
                required
                className="block w-full rounded-md ring-1 ring-gray-300 py-2 px-3 text-gray-900 shadow-sm focus:ring-blue-600 focus:outline-none"
              />
            </div>
          </div>

          <div>
            <button
              type="submit"
              className="flex w-full justify-center rounded-md bg-blue-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm enabled:hover:bg-blue-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-blue-600 disabled:opacity-50"
              disabled={!email || !password}
            >
              Sign in
            </button>
          </div>
        </form>

        <p className="mt-10 text-center text-sm text-gray-500">
          Already registered?{" "}
          <Link
            href="/auth/login"
            className="font-semibold leading-6 text-blue-600 hover:text-blue-500"
          >
            Click here to login
          </Link>
        </p>
      </div>
    </div>
  );
}
