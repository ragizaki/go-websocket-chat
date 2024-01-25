import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import AuthContextProvider from "@/context/AuthContext";
import WebSocketContextProvider from "@/context/WebSocketContext";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Chatr",
  description:
    "Real-time chat application built with Go (Gin router) and Next.js",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <AuthContextProvider>
      <WebSocketContextProvider>
        <html lang="en">
          <body className={inter.className}>
            <div className="flex flex-col md:flex-row h-full min-h-screen p-5">
              {children}
            </div>
          </body>
        </html>
      </WebSocketContextProvider>
    </AuthContextProvider>
  );
}
