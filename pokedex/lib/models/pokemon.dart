import 'package:flutter/material.dart';

class Pokemon {
  final String name;
  final int number;
  final String imageUrl;
  final String description;
  final Color backgroundColor;

  Pokemon({
    required this.name,
    required this.number,
    required this.imageUrl,
    required this.description,
    required this.backgroundColor,
  });
}

final List<Pokemon> pokemonList = [
  Pokemon(
    name: 'Bulbasaur',
    number: 1,
    imageUrl: 'assets/images/bulbasaur.png',
    description:
        'A strange seed was planted on its back at birth. The plant sprouts and grows with this Pokémon.',
    backgroundColor: const Color(0xFFE8F3E8),
  ),
  Pokemon(
    name: 'Ivysaur',
    number: 2,
    imageUrl: 'assets/images/ivysaur.png',
    description:
        'When the bulb on its back grows large, it appears to lose the ability to stand on its hind legs.',
    backgroundColor: const Color(0xFFE8F3E8),
  ),
  Pokemon(
    name: 'Venusaur',
    number: 3,
    imageUrl: 'assets/images/venusaur.png',
    description:
        'In order to support its flower, which has grown larger due to Mega Evolution, its back and legs have become stronger.',
    backgroundColor: const Color(0xFFE8F3E8),
  ),
  Pokemon(
    name: 'Charmander',
    number: 4,
    imageUrl: 'assets/images/charmander.png',
    description:
        'The flame on its tail indicates its life force. If it is healthy, the flame burns brightly.',
    backgroundColor: const Color(0xFFFCECE7),
  ),
];
