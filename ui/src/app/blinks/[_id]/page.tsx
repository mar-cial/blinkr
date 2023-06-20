import Link from "next/link";
import { Blink } from "../page";

async function getBlink(_id: string) {
  const res = await fetch(`http://rest-server:8000/blinks/list/${_id}`);
  return res.json();
}

export default async function SingleBlink({
  params: { _id },
}: {
  params: { _id: string };
}) {
  const blink: Blink = await getBlink(_id);

  return (
    <>
      <div>
        <Link href={'/blinks'}>
          {`< back`}
        </Link>
        <div className="flex flex-col gap-4">
          <div>
            <h2 className="text-2xl">{blink.title}</h2>
            <p className="text-sm">ID: {blink._id}</p>
          </div>
          <div>
            {blink.message}
          </div>
        </div>
      </div>

    </>
  );
}
