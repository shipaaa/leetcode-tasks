# Task Description. [First Unique Character in a String](https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/881/)

## RUS

Задана строка `s`, найдите в ней первый неповторяющийся символ и верните его индекс. Если он не существует, верните `-1`.

## ENG

Given a string `s`, find the first non-repeating character in it and return its index. If it does not exist, return `-1`.

## Объяснение
1. Создаем словарь `charCount`, где будем считать количество символов в исходной строке `s` 
2. В цикле посчитали сколько раз встречается каждый символ
3. В новом цикле проходимся по строке с самого начала (с самого первого символа)
4. Если какой-то символ встречается только один раз, то возвращаем его индекс
5. В остальных случаях возвращаем `-1`: если строка пустая или если в строке все символы встречаются > 1 раза