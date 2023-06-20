'use client'

import { useState } from "react"
import { Blink } from "./page"

const postBlink = async (data: Blink) => {

    const res = await fetch("http://localhost:8000/blinks/create/one", {
        method: "POST",
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(data)
    })

    if (!res.ok) {
        throw new Error("failed")
    }

    return res.json()
}

export default function InsertBlink() {
    const [blink, setBlink] = useState<Blink>({ title: "", message: "" })

    const handleInput = (e: React.FormEvent) => {
        e.preventDefault()

        postBlink(blink)

        setBlink({title: "", message: ""})
    }

    return (
        <>
            <form onSubmit={handleInput} className="flex flex-col gap-4">
                <div className="flex flex-col gap-2 ">
                    <label htmlFor="title">Title</label>
                    <input type="text" onChange={(e) => { setBlink({ ...blink, title: e.target.value }) }} id="title" name="title" value={blink.title} required className="bg-transparent border-zinc-300 rounded-md border-2 border-opacity-25" />
                </div>

                <div className="flex flex-col gap-2 ">
                    <label htmlFor="message">Message</label>
                    <input type="text" id="message" name="message" value={blink.message} onChange={(e) => { setBlink({ ...blink, message: e.target.value }) }} required className="bg-transparent border-zinc-300 rounded-md border-2 border-opacity-25" />
                </div>

                <div className="flex flex-col">
                    <button type="submit" className="bg-opacity-30 bg-emerald-600 text-emerald-200 py-2 rounded-md">Create</button>
                </div>
            </form>

        </>
    )

}
