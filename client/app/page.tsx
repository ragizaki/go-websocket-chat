"use client";

import { useState, useEffect, useContext } from "react";
import { AuthContext } from "@/context/AuthContext";
import { WebSocketContext } from "@/context/WebSocketContext";
import Rooms from "@/components/Rooms";
import { API_URL, WS_URL } from "@/constants";
import { useRouter } from "next/navigation";

export interface Room {
  id: string;
  name: string;
  owner: string;
}

export default function Home() {
  const [roomName, setRoomName] = useState("");
  const [rooms, setRooms] = useState<Room[]>([]);
  const { user } = useContext(AuthContext);
  const { setConnection } = useContext(WebSocketContext);
  const router = useRouter();

  const getRooms = async () => {
    const res = await fetch(`${API_URL}/api/ws/rooms`, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
    const rooms = await res.json();
    setRooms(rooms);
  };

  const createRoom = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const res = await fetch(`${API_URL}/api/ws/rooms`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      credentials: "include",
      body: JSON.stringify({ name: roomName, owner: user.username }),
    });

    setRoomName("");

    if (!res.ok) {
      if (res.status === 400) {
        throw new Error("Room name already exists");
      }
      throw new Error("Unable to create room");
    }

    const newRoom = await res.json();
    setRooms((prevRooms) => [...prevRooms, newRoom]);
  };

  const joinRoom = async (roomId: string) => {
    const { id, username } = user;
    const ws = new WebSocket(
      `${WS_URL}/api/ws/rooms/${roomId}/join?userId=${id}&username=${username}`
    );
    ws.onopen = () => {
      setConnection(ws);
      return router.push(`/room`);
    };

    ws.onerror = (err) => {
      alert("Could not join room, websocket error");
      console.error(err);
    };
  };

  useEffect(() => {
    getRooms();
  }, []);

  return (
    <main className="my-8 px-4 md:mx-10 w-full h-full">
      <form
        className="flex justify-center my-5 p-3 space-x-3"
        onSubmit={createRoom}
      >
        <input
          value={roomName}
          onChange={(e) => setRoomName(e.target.value)}
          type="text"
          className="p-3 rounded-md border-2 border-gray-300 outline-none focus:border-blue-400"
          placeholder="Room Name"
        />
        <button
          type="submit"
          className="bg-blue-500 p-3 rounded-md text-white enabled:hover:bg-blue-600 transition ease-in disabled:bg-blue-200"
          disabled={!roomName}
        >
          Create Room
        </button>
      </form>
      <Rooms rooms={rooms} joinRoom={joinRoom} />
    </main>
  );
}
