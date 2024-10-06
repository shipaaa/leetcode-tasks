# Task Description. [Single Number](https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/549/)

## RUS

Дан непустой массив целых чисел `nums`, каждый элемент появляется дважды, за исключением одного. Найдите этот единственный элемент.

Необходимо реализовать решение с линейной сложностью во время выполнения и использовать только постоянное дополнительное пространство.

## ENG

Given a **non-empty** array of integers `nums`, every element appears _twice_ except for one. Find that single one.
You must implement a solution with a linear runtime complexity and use only constant extra space.

## Объяснение
1. Создаем промежуточную мапу, ключом которой является значение из нашего слайса. А значение количество раз, которое оно встречается
2. В цикле проходимся по искомому слайсу и увеличиваем счетчик у каждого ключа. В конечном итоге получаем мапу из значений и их повторов в слайсе
3. Итерируемся по мапе и если какое-то значение встречается всего 1 раз, то возвращаем его и выходим из цикла