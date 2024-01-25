import { Room } from "@/app/page";
import React from "react";

interface Props {
  rooms: Room[];
  joinRoom: (roomId: string) => void;
}

export default function Rooms({ rooms, joinRoom }: Props) {
  return (
    <div className="grid grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-5">
      {rooms.map(({ id, name, owner }) => (
        <div
          key={id}
          className="border-2 border-blue-400 p-4 flex items-center rounded-md w-full"
        >
          <div className="w-full">
            <p className="text-sm">{owner}'s room</p>
            <p className="text-blue text-xl font-bold">{name}</p>
          </div>
          <button
            className="rounded bg-blue-500 text-white px-2 py-1 hover:bg-blue-600 transition ease-in"
            onClick={() => joinRoom(id)}
          >
            Join
          </button>
        </div>
      ))}
    </div>
  );
}
