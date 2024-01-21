import React from "react";

interface CardComponent {
  id: number;
  name: string;
  email: string;
}

export default function CardComponent({ email, name }: CardComponent) {
  return (
    <div>
      <h2>{name}</h2>
      <p>{email}</p>
    </div>
  );
}

