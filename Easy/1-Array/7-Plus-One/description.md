# Task Description. [Plus One](https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/559/)

## RUS

Вам дано **большое целое число**, представленное в виде целого массива `digits`, где каждая `digits[i]` является `i-й` цифрой целого числа. Цифры расположены слева направо в порядке от наиболее значимого к наименее значимому. Большое целое число не содержит начальных 0.

Увеличьте большое целое число на единицу и верните результирующий массив цифр.

## ENG

You are given a **large integer** represented as an integer array `digits`, where each `digits[i]` is the `i'th` digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading `0`'s.

Increment the large integer by one and return _the resulting array of digits_.
## Объяснение
1. Проходимся в обратном порядке по искомому слайсу
2. Если последнее число не 9-ка, то просто инкрементируем это число и возвращаем слайс
3. Если последнее число — 9, то превращаем ее в 0 и дальше идем по слайсу. Если еще раз встречается 9, то проделываем ту же операцию
4. Если наш слайс закончился и теперь там все нули, то добавляем единицу в начало (как пример: `[9,9,9]`)
