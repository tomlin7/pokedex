import 'package:flutter/material.dart';
import 'package:pokedex/models/pokemon.dart';
import 'package:pokedex/services/api.dart';

class PokemonProvider extends ChangeNotifier {
  final ApiService _apiService = ApiService();
  List<Pokemon> _pokemonList = [];
  List<Pokemon> _filteredList = [];
  bool _isLoading = false;
  String _error = '';

  List<Pokemon> get pokemonList =>
      _filteredList.isEmpty ? _pokemonList : _filteredList;
  bool get isLoading => _isLoading;
  String get error => _error;

  Future<void> loadPokemon() async {
    _isLoading = true;
    _error = '';
    notifyListeners();

    try {
      _pokemonList = await _apiService.getAllPokemon();
      _filteredList = [];
    } catch (e) {
      _error = 'Failed to load Pokemon';
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }

  Future<void> searchPokemon(String query) async {
    if (query.isEmpty) {
      _filteredList = [];
      notifyListeners();
      return;
    }

    _isLoading = true;
    notifyListeners();

    try {
      _filteredList = await _apiService.searchPokemon(query);
    } catch (e) {
      _error = 'Search failed';
    } finally {
      _isLoading = false;
      notifyListeners();
    }
  }
}
