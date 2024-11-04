class Item{
  final int id;
  final String title;
  final String image;
  final String description;
  final String rules;
  final int age;
  final String gamers;
  final String time;
  final int price;
  final int indicator;
  bool isFavorite;

  Item(
      {required this.id,
        required this.title,
        required this.image,
        required this.description,
        required this.rules,
        required this.age,
        required this.gamers,
        required this.time,
        required this.price,
        required this.indicator,
        required this.isFavorite
      });

  factory Item.fromJson(Map<String, dynamic> json) {
    return Item(
      id: json['ID'].toInt() ?? 0,
      title: json['Title'] ?? '',
      image: json['ImageURL'] ?? '',
      description: json['Description'] ?? '',
      rules: json['Rules'] ?? '',
      age: json['Age'].toInt() ?? 0,
      price: json['Price'].toInt() ?? 0,
      gamers: json['Gamers'] ?? '',
      time: json['Time'] ?? '',
      indicator: json['Indicator'] ?? 1,
      isFavorite: json['IsFavorite'] ?? false
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'ID': id,
      'Title': title,
      'ImageURL': image,
      'Description': description,
      'Rules': rules,
      'Age': age,
      'Price': price,
      'Gamers': gamers,
      'Time': time,
      'Indicator': indicator
    };
  }
}

class BasketItem{
  final int id;
  int counter;

  BasketItem(
    {
      required this.id,
      required this.counter
    }
  );

  factory BasketItem.fromJson(Map<String, dynamic> json) {
    return BasketItem(
      id: json['ID'].toInt() ?? 0,
      counter: json['Counter'] ?? 0,
    );
  }
}

class User{
  String name;
  String surname;
  String patronymic;
  String email;
  String telNumber;
  String image;

  User(this.name, this.surname, this.patronymic, this.email, this.telNumber, this.image);
}

User admin = User('Анастасия', 'Полошкова', 'Юрьевна', 'poloshkova.a.y@edu.mirea.ru', '+7 (999) 999-99-99', 'https://i.pinimg.com/474x/ed/b6/d9/edb6d911b0edf65204fb3751c61c5fa9.jpg');