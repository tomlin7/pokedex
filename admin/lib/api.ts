import axios from "axios";
import { Pokemon } from "./pokemon";

const API_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export async function getAllPokemon(): Promise<Pokemon[]> {
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

export async function createPokemon(pokemon: Pokemon): Promise<Pokemon> {
  const res = await axios.post(`${API_URL}/api/pokemon`, pokemon);
  if (res.status !== 201) {
    throw new Error("Failed to create pokemon");
  }

  return res.data;
}

export async function updatePokemon(
  id: string,
  pokemon: Pokemon
): Promise<Pokemon> {
  const res = await axios.put(`${API_URL}/api/pokemon/${id}`, pokemon);
  if (res.status !== 200) {
    throw new Error("Failed to update pokemon");
  }

  return res.data;
}

export async function deletePokemon(id: string): Promise<void> {
  const res = await axios.delete(`${API_URL}/api/pokemon/${id}`);
  if (res.status !== 200) {
    throw new Error("Failed to delete pokemon");
  }
}
