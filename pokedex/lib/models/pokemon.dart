import 'package:flutter/material.dart';

class Pokemon {
  final String id;
  final String name;
  final int number;
  final String imageUrl;
  final String description;
  final Color backgroundColor;

  Pokemon({
    required this.id,
    required this.name,
    required this.number,
    required this.imageUrl,
    required this.description,
    required this.backgroundColor,
  });

  factory Pokemon.fromJson(Map<String, dynamic> json) {
    return Pokemon(
      id: json['id'],
      name: json['name'],
      number: json['number'],
      imageUrl: json['imageUrl'],
      description: json['description'],
      backgroundColor: _hexToColor(json['backgroundColor']),
    );
  }

  static Color _hexToColor(String hex) {
    hex = hex.replaceAll('#', '');
    return Color(int.parse('FF$hex', radix: 16));
  }
}
