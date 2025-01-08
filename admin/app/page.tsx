import PokemonList from "@/components/list";
import { Suspense } from "react";

export default function Home() {
  return (
    <div className="container mx-auto py-8 px-4">
      <Suspense fallback={<div>Loading...</div>}>
        <PokemonList />
      </Suspense>
    </div>
  );
}
