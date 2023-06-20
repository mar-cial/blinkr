import Link from "next/link";
import InsertBlink from "./insert";

async function getData() {
  try {
    const res = await fetch("http://rest-server:8000/blinks/list", { cache: "no-store" })
    if (!res.ok) {
      throw new Error("Network response was not ok")
    }

    return res.json()
  } catch (err) {
    console.log("There was a problem with your operation.")
    console.log(err)
    return []
  }
}

export interface Blink {
  _id?: string;
  title: string;
  message: string;
}

export default async function Home() {
  const data: Blink[] = await getData();

  return (
    <>
      <div>
        <h3 className="text-3xl font-semibold">Insert new</h3>
        <InsertBlink />
      </div>

      <div>
        <h3 className="text-3xl font-semibold">Current blinks</h3>
      </div>


      <div>
        {!data ? (
          <div>
            <h3>Empty!</h3>
          </div>
        ) : (
          <div className={'grid gap-2 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4'}>
            {
              data.map((b) => {
                return (
                  <Link href={`/blinks/${b._id}`} key={b._id}>
                    <div key={b._id}>
                      <h3 className="text-xl font-medium">{b.title}</h3>
                    </div>
                  </Link>
                )
              })
            }
          </div>
        )}
      </div>
    </>
  );
}
