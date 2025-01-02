import 'package:flutter/material.dart';
import 'package:pokedex/providers/pokemon.dart';
import 'package:pokedex/widgets/grid.dart';
import 'package:pokedex/widgets/search.dart';
import 'package:provider/provider.dart';

class Home extends StatelessWidget {
  const Home({super.key});

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // appBar: AppBar(
      //   backgroundColor: Colors.transparent,
      //   elevation: 0,
      //   actions: [
      //     IconButton(
      //       icon: const Icon(Icons.menu),
      //       onPressed: () {},
      //     ),
      //   ],
      // ),
      body: SafeArea(
        child: Padding(
          padding: const EdgeInsets.only(left: 20, right: 20, top: 20),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text(
                'Pokédex',
                style: TextStyle(
                  fontSize: 32,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF2E3057),
                ),
              ),
              const SizedBox(height: 8),
              const Text(
                'Search for a Pokémon by name or using its\nNational Pokédex number.',
                style: TextStyle(
                  color: Colors.grey,
                  fontSize: 14,
                ),
              ),
              const SizedBox(height: 20),
              const CustomSearchBar(),
              const SizedBox(height: 20),
              Expanded(
                child: Consumer<PokemonProvider>(
                  builder: (context, provider, _) {
                    if (provider.isLoading) {
                      return const Center(child: CircularProgressIndicator());
                    }
                    if (provider.error.isNotEmpty) {
                      return Center(child: Text(provider.error));
                    }
                    return const PokemonGrid();
                  },
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
