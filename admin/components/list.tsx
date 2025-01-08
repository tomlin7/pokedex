"use client";

import { getAllPokemon, Pokemon } from "@/lib/api";
import { useEffect, useState } from "react";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";

export default function PokemonList() {
  const [pokemon, setPokemon] = useState<Pokemon[]>([]);

  useEffect(() => {
    getAllPokemon().then(setPokemon);
  }, []);

  return (
    <>
      <div className="rounded-md border">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead>Number</TableHead>
              <TableHead>Image</TableHead>
              <TableHead>Name</TableHead>
              <TableHead>Description</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {pokemon.map((p) => (
              <TableRow key={p.id}>
                <TableCell>{p.number}</TableCell>
                <TableCell>
                  <img
                    src={p.imageUrl}
                    alt={p.name}
                    className="w-12 h-12 object-contain"
                    style={{ backgroundColor: p.backgroundColor }}
                  />
                </TableCell>
                <TableCell>{p.name}</TableCell>
                <TableCell className="max-w-md truncate">
                  {p.description}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>
    </>
  );
}
