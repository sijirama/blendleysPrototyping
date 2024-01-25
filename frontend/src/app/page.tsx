"use client"

import { useRouter } from "next/navigation";

export default function Home() {
    const route = useRouter()
    const onClick = () => {
        route.push("/add")
    }
    return (
        <main className="flex min-h-screen flex-col items-center justify-center p-24">
            Hello world
            <button onClick={onClick} className="py-2 px-5 bg-zinc-950 text-zinc-300 rounded-md">Add a part.</button>
        </main>
    );
}
