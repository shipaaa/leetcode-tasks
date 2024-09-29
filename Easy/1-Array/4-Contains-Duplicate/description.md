# Task Description. [Contains Duplicate](https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/578/)

## RUS

Вам дан целочисленный массив `nums`, верните `true`, если какое-либо значение появляется в массиве по крайней мере дважды, и верните `false`, если каждый элемент различен.

## ENG

Given an integer array `nums`, return `true` if any value appears **at least twice** in the array, and return `false` if every element is distinct.
## Объяснение
1. Создаем `map`, в которой инкрементируем значение у ключа на каждом проходе цикла. Ключом является значение из `nums`
2. Если у какого-то ключа значение становится больше 1 (`> 1`) — возвращаем `true`
3. Проходимся по всему списку и если ничего не находим, возвращаем `false`