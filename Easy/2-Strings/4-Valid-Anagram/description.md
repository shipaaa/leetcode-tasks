# Task Description. [Valid Anagram](https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/882/)

## RUS

При наличии двух строк `s` и `t` верните значение true, если `t` является анаграммой `s`, и значение false в противном случае.

Анаграмма — это слово или фраза, образованные путем перестановки букв другого слова или фразы, обычно с использованием всех исходных букв ровно один раз.

## ENG

Given two strings `s` and `t`, return `true` _if_ `t` _is an anagram of_ `s`_, and_ `false` _otherwise_.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

## Объяснение
1. Создаем 2 слайса `[]int` для того чтобы положить туда unicode значение каждого символа в строках `s` и `t`. Делаем это в цикле
2. Отсортировываем 2 слайса
3. Проходимся в цикле по этим слайсам и сравниваем каждое значение: Если хотя бы одно значение из 1-го слайса не равно значению из 2-го слайса, то возвращаем `false`. Если все ок — `true`