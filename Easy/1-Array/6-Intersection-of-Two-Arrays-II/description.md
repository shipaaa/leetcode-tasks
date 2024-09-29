# Task Description. [Intersection of Two Arrays II](https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/674/)

## RUS

Дано два целочисленных массива `nums1` и `nums2`, верните массив их пересечений. Каждый элемент в результате должен появляться столько раз, сколько он отображается в обоих массивах, и вы можете возвращать результат в **любом порядке**.

## ENG

Given two integer arrays `nums1` and `nums2`, return _an array of their intersection_. Each element in the result must appear as many times as it shows in both arrays and you may return the result in **any order**.
## Объяснение
1. Создаем `map`, в которой инкрементируем значение у ключа на каждой итерации цикла для `nums1`. Ключом является значение из `nums1`
2. Запускаем новый цикл, в котором итерируемся по `nums2`. Если значение в конкретной итерации содержится в `map` и значение (счетчик) этого ключа > 0, то добавляем значение в конечный slice `intersection`. Декрементируем счетчик после добавления в slice