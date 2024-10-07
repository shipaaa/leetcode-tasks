# Task Description. [Rotate Array](https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/646/)

## RUS

Дан целочисленный массив `nums`, поверните массив вправо на `k` шагов, где `k` неотрицательно.

## ENG

Given an integer array `nums`, rotate the array to the right by `k` steps, where `k` is non-negative.
## Объяснение
1. `k` будет равно остатку от деления числа `k` на длину слайса
2. Создаем функцию `reverse`, которая будет менять местами левый и правый элемент в слайсе, который мы передаем
3. Таким образом, на 1 шаге разворачиваем исходный слайс полностью
4. После этого, разворачиваем первые `k` элементов измененного слайса
5. Дальше разворачиваем оставшиеся после `k` элементы