# Task Description. [Two Sum](https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/546/)

## RUS

Дан массив целых чисел `nums` и int `target`, верните индексы двух чисел таким образом, чтобы они в сумме были равны `target`.

Вы можете предположить, что у каждого входного сигнала будет ровно одно решение, и вы не сможете использовать один и тот же элемент дважды.

Вы можете вернуть ответ в любом порядке.

## ENG

Given an array of integers `nums` and an integer `target`, return _indices of the two numbers such that they add up to `target`_.

You may assume that each input would have **_exactly_ one solution**, and you may not use the _same_ element twice.

You can return the answer in any order.

## Объяснение
1. Создаем промежуточную карту `interimM`, ключом которой является значение элемента из слайса , а значением — является индекс этого элемента из слайса
2. В цикле проходимся по искомому слайсу
3. Задаем переменную `requiredValue`. Она равна разности `target` и числа на каждой итерации
4. Проверяем существует ли ключ `requiredValue` в промежуточной карте. Если да, то возвращаем индексы элементов: значение из мапы по ключу и индекс слайса на текущей итерации. Если нет, то обновляем мапу, добавляя новый ключ-значение