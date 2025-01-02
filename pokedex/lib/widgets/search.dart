import 'package:flutter/material.dart';
import 'package:pokedex/providers/pokemon.dart';
import 'package:pokedex/services/debouncer.dart';
import 'package:provider/provider.dart';

class CustomSearchBar extends StatefulWidget {
  const CustomSearchBar({super.key});

  @override
  State<CustomSearchBar> createState() => _CustomSearchBarState();
}

class _CustomSearchBarState extends State<CustomSearchBar> {
  final _debouncer = Debouncer(milliseconds: 500);
  final _searchController = TextEditingController();

  void _onSearchChanged(String query) {
    _debouncer.run(() {
      context.read<PokemonProvider>().searchPokemon(query);
    });
  }

  @override
  void dispose() {
    _debouncer.dispose();
    _searchController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 16),
      decoration: BoxDecoration(
        color: Colors.grey[200],
        borderRadius: BorderRadius.circular(12),
      ),
      child: Row(
        children: [
          const Icon(Icons.search, color: Colors.grey),
          const SizedBox(width: 10),
          Expanded(
            child: TextField(
              controller: _searchController,
              onChanged: _onSearchChanged,
              decoration: const InputDecoration(
                hintText: 'Search by name or number',
                border: InputBorder.none,
                hintStyle: TextStyle(color: Colors.grey),
              ),
            ),
          ),
          IconButton(
            icon: const Icon(Icons.clear),
            onPressed: () {
              _searchController.clear();
              _onSearchChanged('');
            },
          )

          // Container(
          //   padding: const EdgeInsets.all(8),
          //   decoration: BoxDecoration(
          //     color: const Color(0xFF2E3057),
          //     borderRadius: BorderRadius.circular(8),
          //   ),
          //   child: const Icon(
          //     Icons.tune,
          //     color: Colors.white,
          //     size: 20,
          //   ),
          // )
        ],
      ),
    );
  }
}
