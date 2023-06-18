import { Azeret_Mono } from "next/font/google";
import Link from "next/link";

const main = Azeret_Mono({ subsets: ["latin"] });

export default function Home() {
	return (
    <div className="flex flex-col gap-12 p-24">
      <h1 className={`text-4xl ${main.className}`}>Blinkr.</h1>

	  <Link href={'/blinks'}>blinks.</Link>
    </div>
	)
}