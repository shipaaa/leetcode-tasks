# Task Description. [Valid Sudoku](https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/769/)

## RUS

Определите, подходит ли доска для судоку размером `9 х 9`. Только ячейки поля должны быть проверены в соответствии со следующими правилами:
1. Каждая строка должна содержать цифры `от 1 до 9` без повторений.
2. Каждый столбец должен содержать цифры от `1 до 9` без повторений.
3. Каждая из девяти ячеек сетки размером `3 х 3` должна содержать цифры `от 1 до 9` без повторений.
   
**Примечание:**
- Доска для судоку (частично заполненная) может быть действительной, но не обязательно разрешимой.
- Только заполненные ячейки должны быть проверены в соответствии с указанными правилами.

## ENG

Determine if a `9 x 9` Sudoku board is valid. Only the filled cells need to be validated **according to the following rules**:

1. Each row must contain the digits `1-9` without repetition.
2. Each column must contain the digits `1-9` without repetition.
3. Each of the nine `3 x 3` sub-boxes of the grid must contain the digits `1-9` without repetition.
   
**Note:**
- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

## Объяснение
1. Создаем 3 массива по 9 элементов каждый, состоящих из мап, для отслеживания встречающихся чисел в строках, столбцах и блоках 3x3
2. Проходим по каждой ячейке доски. Если ячейка пустая (в нашем случае = `.`), то идем дальше
3. Если число не встречалось раньше, мы его запоминаем, устанавливая соответствующее значение в мапе `true`:
4. Если же число уже было в какой-то мапе выводим `false`