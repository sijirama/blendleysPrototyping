"use client"
import { useState } from "react";
import { useRouter } from "next/navigation";
import Editor from "@/components/Editor2";

export default function Home() {
    const [data, setData] = useState();
    const route = useRouter()

    const onClick = () => {
        route.push("/add")
    }

    const onSubmit = () => {
        console.log(data)
    }

    return (
        <div className="grid grid-cols-3 gap-2 m-5 p-3">
            <div className="col-span-2 w-full">
                <h1 className="my-5 font-bold">Blendle</h1>
                    <div className="border rounded-md p-5">
                        <Editor />
                    </div>
                    <button form="community-post-form" onClick={onSubmit} className=" mt-5 py-1 px-3 text-sm rounded-md text-zinc-100 hover:bg-zinc-700 bg-zinc-900" type="submit">Submit</button>
                    <button onClick={onClick} className=" mx-4 mt-5 py-1 px-3 text-sm rounded-md text-zinc-100 hover:bg-zinc-700 bg-zinc-900" type="submit">Go Home</button>
            </div>
        </div>
    );
}
