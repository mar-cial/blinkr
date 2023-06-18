import { Azeret_Mono } from "next/font/google";
import Link from "next/link";

const main = Azeret_Mono({ subsets: ["latin"] });

const getData = async () => {
  const res = await fetch("http://localhost:8000/blinks/list");

  if (!res.ok) {
    throw new Error("Failed to fetch data");
  }

  return res.json();
};

export interface Blink {
  _id: string;
  title: string;
  message: string;
}

export default async function Home() {
  const data: Blink[] = await getData();

  return (
    <div className="flex flex-col gap-12 p-24">
      <h1 className={`text-4xl ${main.className}`}>Blinkr.</h1>

      <div className="flex flex-col gap-6">
        {data.map((b) => (
          <Link key={b._id} href={`/blinks/${b._id}`}>
            <div>
              <h3>{b.title}</h3>
              <p>{b.message.slice(0, 20)}...</p>
            </div>
          </Link>
        ))}
      </div>
    </div>
  );
}
