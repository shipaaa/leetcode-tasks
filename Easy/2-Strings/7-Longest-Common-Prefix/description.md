# Task Description. [Longest Common Prefix](https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/887/)

## RUS

Напишите функцию, которая найдет самую длинную строку с общим префиксом среди массива строк.
Если общего префикса нет, верните пустую строку `""`.

## ENG

Write a function to find the longest common prefix string amongst an array of strings.
If there is no common prefix, return an empty string `""`.

## Объяснение
1. Предположим что первая строка в слайсе это и есть наш общий префикс. `prefix`
2. В цикле проходим по всем строкам, начиная со второго элемента слайса
3. Создаем еще один внутренний цикл и выполняем его до тех пор, пока текущая строка `strs[i]` не начинается с `prefix` и длина `prefix` > 0. Если не начинается: `prefix` постепенно сокращается, пока не будет найдено совпадение
4. Функция `startswith` используется для проверки, начинается ли строка `s` с заданного префикса. 