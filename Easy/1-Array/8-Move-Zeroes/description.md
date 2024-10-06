# Task Description. [Move Zeroes](https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/567/)

## RUS

Дан целочисленный массив `nums`, переместите все `0` в его конец, сохраняя относительный порядок ненулевых элементов.

**Обратите внимание**, что вы должны сделать это на месте, не создавая копию массива.

## ENG

Given an integer array `nums`, move all `0`'s to the end of it while maintaining the relative order of the non-zero elements.

**Note** that you must do this in-place without making a copy of the array.
## Объяснение
1. Объявляем переменную `lastNonZero`, которая отвечает за индекс последнего ненулевого эл-та в слайсе
2. В цикле поочередно проверяем является ли текущий элемент ненулевым. Если является, то перемещаем этот элемент на позицию `lastNonZero`. Инкрементируем `lastNonZero`, чтобы указать на следующую свободную позицию для ненулевого элемента.
3. После полного прохода слайса: в новом цикле заполняем оставшуюся часть слайса нулями начиная с индекса `lastNonZero` до конца слайса