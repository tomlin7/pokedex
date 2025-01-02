import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:pokedex/models/pokemon.dart';

class ApiService {
  static const String baseUrl = 'https://pokedex-n655.onrender.com/api';

  Future<List<Pokemon>> getAllPokemon() async {
    final response = await http.get(Uri.parse('$baseUrl/pokemon'));

    if (response.statusCode == 200) {
      final List<dynamic> jsonList = json.decode(response.body);
      return jsonList.map((json) => Pokemon.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load pokemon');
    }
  }

  Future<Pokemon> getPokemonById(int id) async {
    final response = await http.get(Uri.parse('$baseUrl/pokemon/$id'));

    if (response.statusCode == 200) {
      return Pokemon.fromJson(json.decode(response.body));
    } else {
      throw Exception('Failed to load pokemon details');
    }
  }

  Future<List<Pokemon>> searchPokemon(String query) async {
    final response = await http.get(
      Uri.parse('$baseUrl/pokemon/search').replace(
        queryParameters: {'q': query},
      ),
    );

    if (response.statusCode == 200) {
      final List<dynamic> jsonList = json.decode(response.body);
      return jsonList.map((json) => Pokemon.fromJson(json)).toList();
    } else {
      throw Exception('Failed to search pokemon');
    }
  }
}
