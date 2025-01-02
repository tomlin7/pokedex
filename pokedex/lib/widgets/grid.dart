import 'package:flutter/material.dart';
import 'package:pokedex/providers/pokemon.dart';
import 'package:pokedex/screens/detail.dart';
import 'package:pokedex/widgets/card.dart';
import 'package:provider/provider.dart';

class PokemonGrid extends StatelessWidget {
  const PokemonGrid({super.key});

  @override
  Widget build(BuildContext context) {
    final pokemonList = context.watch<PokemonProvider>().pokemonList;

    return GridView.builder(
      gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
        crossAxisCount: 2,
        childAspectRatio: 0.9,
        crossAxisSpacing: 10,
        mainAxisSpacing: 10,
      ),
      itemCount: pokemonList.length,
      itemBuilder: (context, index) {
        return GestureDetector(
          onTap: () {
            Navigator.push(
              context,
              MaterialPageRoute(
                builder: (context) => DetailScreen(
                  pokemon: pokemonList[index],
                ),
              ),
            );
          },
          child: PokemonCard(pokemon: pokemonList[index]),
        );
      },
    );
  }
}
