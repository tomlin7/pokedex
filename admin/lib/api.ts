import axios from "axios";
import { z } from "zod";

export const PokemonSchema = z.object({
  id: z.string().optional(),
  name: z.string().min(1, "Name is required"),
  number: z.string().min(1, "Number must be positive"),
  imageUrl: z.string().url("Must be a valid URL"),
  description: z.string().min(1, "Description is required"),
  backgroundColor: z
    .string()
    .regex(/^([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$/, "Must be a valid hex color"),
});

export type Pokemon = z.infer<typeof PokemonSchema>;

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export async function getAllPokemon(): Promise<Pokemon[]> {
  console.log(`${API_URL}/api/pokemon`);
  const res = await axios.get(`${API_URL}/api/pokemon`);
  if (res.status !== 200) {
    throw new Error(`Failed to fetch pokemon ${res.status}`);
  }

  return res.data;
}

export async function getPokemon(id: string): Promise<Pokemon> {
  const res = await axios.get(`${API_URL}/api/pokemon/${id}`);
  if (res.status !== 200) {
    throw new Error("Failed to fetch pokemon");
  }

  return res.data;
}

export async function searchPokemon(query: string): Promise<Pokemon[]> {
  const res = await axios.get(
    `${API_URL}/api/pokemon/search?q=${encodeURIComponent(query)}`
  );
  if (res.status !== 200) {
    throw new Error("Failed to search pokemon");
  }

  return res.data;
}
