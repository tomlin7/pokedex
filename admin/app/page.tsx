"use client";

import { PokemonDialog } from "@/components/dialog";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import {
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useToast } from "@/hooks/use-toast";
import { getAllPokemon, searchPokemon } from "@/lib/api";
import { Pokemon } from "@/lib/pokemon";
import { Search } from "lucide-react";
import { useEffect, useState } from "react";

export default function Home() {
  const [pokemon, setPokemon] = useState<Pokemon[]>([]);
  const [searchQuery, setSearchQuery] = useState("");
  const [isNewDialogOpen, setIsNewDialogOpen] = useState(false);
  const [editingPokemon, setEditingPokemon] = useState<Pokemon | null>(null);
  const { toast } = useToast();

  useEffect(() => {
    loadPokemon();
  }, []);

  async function loadPokemon() {
    try {
      const data = await getAllPokemon();
      setPokemon(data);
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to load Pokemon",
        variant: "destructive",
      });
    }
  }

  async function handleSearch(e: React.FormEvent) {
    e.preventDefault();
    if (!searchQuery.trim()) {
      loadPokemon();
      return;
    }
    try {
      const results = await searchPokemon(searchQuery);
      setPokemon(results);
    } catch (error) {
      toast({
        title: "Error",
        description: "Failed to search Pokemon",
        variant: "destructive",
      });
    }
  }

  return (
    <div className="container mx-auto py-10">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-4xl font-bold">Pokemon</h1>
        <Button onClick={() => setIsNewDialogOpen(true)}>New</Button>
      </div>

      <form onSubmit={handleSearch} className="relative mb-8">
        <Search className="absolute left-3 top-3 h-4 w-4 text-muted-foreground" />
        <Input
          placeholder="Search for a Pokemon"
          value={searchQuery}
          onChange={(e) => setSearchQuery(e.target.value)}
          className="pl-10"
        />
      </form>

      <div className="border rounded-lg">
        <Table>
          <TableHeader>
            <TableRow>
              <TableHead className="w-16">#</TableHead>
              <TableHead>Name</TableHead>
              <TableHead>Description</TableHead>
              <TableHead className="text-right">Actions</TableHead>
            </TableRow>
          </TableHeader>
          <TableBody>
            {pokemon.map((p) => (
              <TableRow key={p.id}>
                <TableCell>{p.number}</TableCell>
                <TableCell>{p.name}</TableCell>
                <TableCell>{p.description}</TableCell>
                <TableCell className="text-right">
                  <Button variant="ghost" onClick={() => setEditingPokemon(p)}>
                    Edit
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </div>

      <PokemonDialog
        open={isNewDialogOpen}
        onOpenChange={setIsNewDialogOpen}
        onSuccess={() => {
          loadPokemon();
          setIsNewDialogOpen(false);
        }}
      />

      {editingPokemon && (
        <PokemonDialog
          open={true}
          onOpenChange={() => setEditingPokemon(null)}
          pokemon={editingPokemon}
          onSuccess={() => {
            loadPokemon();
            setEditingPokemon(null);
          }}
        />
      )}
    </div>
  );
}
