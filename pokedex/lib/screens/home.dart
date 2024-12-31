import 'package:flutter/material.dart';
import 'package:pokedex/widgets/grid.dart';
import 'package:pokedex/widgets/search.dart';

class Home extends StatelessWidget {
  const Home({super.key});

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
        body: SafeArea(
      child: Padding(
        padding: EdgeInsets.all(20.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(
              'Pokédex',
              style: TextStyle(
                  fontSize: 32,
                  fontWeight: FontWeight.bold,
                  color: Color(0xFF2E3057)),
            ),
            SizedBox(height: 8),
            Text(
              'Search for Pokémon by name or using the\nNational Pokédex number.',
              style: TextStyle(
                fontSize: 14,
                color: Colors.grey,
              ),
            ),
            SizedBox(
              height: 20,
            ),
            CustomSearchBar(),
            SizedBox(
              height: 20,
            ),
            Expanded(
              child: PokemonGrid(),
            ),
          ],
        ),
      ),
    ));
  }
}
