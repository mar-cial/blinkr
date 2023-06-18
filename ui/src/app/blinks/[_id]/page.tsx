import { Blink } from "../page";

async function getBlink(_id: string) {
  const res = await fetch(`http://localhost:8000/blinks/list/${_id}`);
  return res.json();
}

export default async function SingleBlink({
  params: { _id },
}: {
  params: { _id: string };
}) {
  const blink: Blink = await getBlink(_id);
  console.log(blink);
  return (
    <div>
      <h2>{blink.title}</h2>
    </div>
  );
}
