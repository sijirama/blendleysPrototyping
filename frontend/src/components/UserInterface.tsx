"use client";
import React, { useEffect, useState } from "react";
import axios from "axios";
import CardComponent from "./CardComponent";

interface User {
  id: number;
  name: string;
  email: string;
}

interface UserInterfaceProps {
  backendName: string;
}

export default function UserInterface({ backendName }: UserInterfaceProps) {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8000";
  const [user, setUser] = useState<User[]>([]);
  const [newUser, setNewUser] = useState({ name: "", email: "" });
  const [updateUser, setUpdateUser] = useState({ id: "", name: "", email: "" });

  const backgroundCOlors: { [key: string]: string } = {
    go: "bg-cyan-500",
  };

  const buttonCOlors: { [key: string]: string } = {
    go: "bg-cyan-700 hover:bg-blue-600",
  };

  useEffect(() => {
    const fetchUsers = async () => {
      try {
        const response = await axios.get(`${apiUrl}/api/${backendName}/users`);
        setUser(response.data.reverse());
      } catch (error) {
        console.log("Error fetching users", error);
      }
    };

    fetchUsers();
  }, [backendName, apiUrl]);

  return (
    <div>
      {user.map((u, i) => (
        <div key={i}>
          <div>{u.name}</div>
          <div>{u.email}</div>
        </div>
      ))}
    </div>
  );
}
