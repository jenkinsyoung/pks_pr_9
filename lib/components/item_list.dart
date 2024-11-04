import 'package:flutter/material.dart';
import 'package:pr_9/components/item_ui.dart';
import 'package:pr_9/models/item.dart';

class ItemList extends StatelessWidget {
  const ItemList({super.key, required this.game});
  final Item game;
  @override
  Widget build(BuildContext context) {
    if (game.indicator == 1) {
      return ItemUi(
        key: Key('${game.id}'),
        game: game,
        bodyColor: const Color.fromRGBO(255, 207, 2, 1.0),
        textColor: const Color.fromRGBO(129, 40, 0, 1.0),
      );
    } else if (game.indicator == 2) {
      return ItemUi(
        key: Key('${game.id}'),
        game: game,
        bodyColor: const Color.fromRGBO(163, 3, 99, 1.0),
        textColor: const Color.fromRGBO(255, 204, 254, 1.0),
      );
    } else {
      return ItemUi(
        key: Key('${game.id}'),
        game: game,
        bodyColor: const Color.fromRGBO(48, 0, 155, 1.0),
        textColor: const Color.fromRGBO(203, 238, 251, 1.0),
      );
    }
  }
}