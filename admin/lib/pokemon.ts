import { z } from "zod";

export const PokemonSchema = z.object({
  id: z.string().optional(),
  name: z.string().min(1, "Name is required"),
  number: z.number().min(1, "Number must be positive"),
  imageUrl: z.string().url("Must be a valid URL"),
  description: z.string().min(1, "Description is required"),
  backgroundColor: z
    .string()
    .regex(/^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$/, "Must be a valid hex color"),
});

export type Pokemon = z.infer<typeof PokemonSchema>;
