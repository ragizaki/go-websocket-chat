"use client";

import { createContext, useState, useEffect } from "react";
import { useRouter } from "next/navigation";

type Connection = WebSocket | null;

export const WebSocketContext = createContext<{
  connection: Connection;
  setConnection: React.Dispatch<React.SetStateAction<Connection>>;
}>({
  connection: null,
  setConnection: () => {},
});

export default function WebSocketContextProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const [connection, setConnection] = useState<Connection>(null);

  return (
    <WebSocketContext.Provider value={{ connection, setConnection }}>
      {children}
    </WebSocketContext.Provider>
  );
}
