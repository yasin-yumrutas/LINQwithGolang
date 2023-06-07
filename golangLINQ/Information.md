{Where, Map, Sort, GroupBy, Join} -> temel LINQ sorguları


Adım 1'in açıklması :
Bu program, bir slice oluşturur ve sonra bu slice üzerinde birkaç LINQ benzeri işlem yapar. İlk olarak, Where işlevi, slice'ın elemanlarını filtrelemek için kullanılır ve sadece 2'den büyük olanları alır. Sonra, Select işlevi, her elemanın değerini 2 ile çarparak yeni bir slice oluşturur. En sonunda, ToSlice işlevi, sonuçları orijinal slice'a geri yazar.


Adım 2'nin açıklaması :
Burda Where işlemlerinin gerçekleştiği bir kod parçası çalıştırdık


Adım 3'ün açıklaması :
Bu sorgu, "numbers" adlı bir slice üzerinde "Map" işlemi yaparak, her elemanın değerini 2 ile çarparak yeni bir slice oluşturur.


Adım 4'ün açıklaması :
Bu sorgu,  "numbers" adlı bir slice içindeki dağınık sırayla yazılmış sayıların küçükten büyüğe doğru sıralanmasını sağladı.(Sort Fonksiyonu)



Adım 5'in açıklaması :
Bu sorgu, "words" adlı bir slice üzerinde "GroupBy" işlemi yaparak, slicedaki kelimeleri ilk harflerine göre gruplar.


Adım 6'nın açıklması :
Bu kod iki farklı veri yapısının (struct) belirli bir anahtara göre birleştirilmesini sağlayan bir fonksiyonel join operasyonu gerçekleştirmektedir.

join fonksiyonu, gelen iki veri yapısını (table1, table2) ve her bir veri yapısı için bir anahtar fonksiyonu (keyFunc1, keyFunc2) alır. Anahtar fonksiyonları, veri yapılarının içindeki bir alanın değerine göre o veri yapısında bulunan verileri birbirine bağlarlar.

Son olarak, bu iki veri yapısında bulunan verileri birleştirmek için resultFunc adlı bir fonksiyon kullanılır. Bu fonksiyon, her bir veri yapısındaki verileri birleştirmek için kullanılır ve sonuçlar result dizisinde birleştirilir.

Sonuç olarak, join fonksiyonu, iki veri yapısını birleştirerek yeni bir veri yapısı oluşturur ve sonucu geri döndürür. Bu kod örneği, join fonksiyonunun nasıl kullanılabileceğini gösterir.