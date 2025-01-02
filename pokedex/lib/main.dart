import 'package:flutter/material.dart';
import 'package:pokedex/providers/pokemon.dart';
import 'package:pokedex/screens/home.dart';
import 'package:provider/provider.dart';

void main() {
  runApp(const PokemonApp());
}

class PokemonApp extends StatelessWidget {
  const PokemonApp({super.key});

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => PokemonProvider()..loadPokemon(),
      child: MaterialApp(
        title: 'Pokédex',
        theme: ThemeData(
          primarySwatch: Colors.blue,
          scaffoldBackgroundColor: const Color(0xFFF5F9FD),
          fontFamily: 'Poppins',
        ),
        home: const Home(),
      ),
    );
  }
}
