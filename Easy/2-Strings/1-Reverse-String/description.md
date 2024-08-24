# Task Description. [Reverse String](https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/879/)

## RUS

Напишите функцию, которая преобразует строку в обратном порядке. Входная строка задается в виде массива символов `s`.

Для этого необходимо изменить входной массив [на месте](https://en.wikipedia.org/wiki/In-place_algorithm), добавив к нему дополнительную память `O(1)`.

## ENG

Write a function that reverses a string. The input string is given as an array of characters `s`.

You must do this by modifying the input array [in-place](https://en.wikipedia.org/wiki/In-place_algorithm) with `O(1)` extra memory.

## Объяснение
1. Задаем переменные `left`, `right` — индекс 1-го и последнего значения в массиве
2. Проходимся по массиву в цикле до тех пор, пока `left` не станет меньше `right`
3. Присваиваем значение в массиве слева значению справа и наоборот. И делаем до тех пор, пока не дойдем до середины массива
4. На каждой итерации: увеличиваем `left`, уменьшаем `right`