# Task Description. [String to Integer (atoi)](https://leetcode.com/explore/interview/card/top-interview-questions-easy/127/strings/884/)

## RUS

Реализуйте функцию `myAtoi(string s)`, которая преобразует строку в 32-разрядное целое число со знаком.

Алгоритм для `myAtoi(string s)` следующий:

1. **Пробел**: игнорируйте все начальные пробелы (`" "`).
2. **Знак**: Определите знак, проверив, является ли следующий символ `-` или `+`, предполагая, что ни один из них не является положительным.
3. **Преобразование**: Считывайте целое число, пропуская начальные нули, до тех пор, пока не встретится нецифровой символ или пока не будет достигнут конец строки. Если цифры не были считаны, то результатом будет 0.
4. **Округление**: Если целое число выходит за пределы диапазона 32-разрядных целых чисел со знаком $[-2^{31}, 2^{31}-1]$, затем округлите целое число, чтобы оно оставалось в пределах диапазона. В частности, целые числа, меньшие, чем $-2^{31}$, должны быть округлены до $-2^{31}$, а целые числа, большие, чем $2^{31}-1$, должны быть округлены до $2^{31}-1$.
Верните целое число в качестве конечного результата.
## ENG

Implement the `myAtoi(string s)` function, which converts a string to a 32-bit signed integer.

The algorithm for `myAtoi(string s)` is as follows:

1. **Whitespace**: Ignore any leading whitespace (`" "`).
2. **Signedness**: Determine the sign by checking if the next character is `'-'` or `'+'`, assuming positivity is neither present.
3. **Conversion**: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
4. **Rounding**: If the integer is out of the 32-bit signed integer range $[-2^{31}, 2^{31}-1]$, then round the integer to remain in the range. Specifically, integers less than $-2^{31}$ should be rounded to $-2^{31}$, and integers greater than $2^{31}-1$ should be rounded to $2^{31}-1$.

Return the integer as the final result.

## Объяснение
1. Объявляем переменные `result`, `startIndex`, `sign` (итоговое значение, индекс с которого начинаем проверять число, знак: принимает значения `1`и `-1`)
2. Пропускаем все пробелы и узнаем знак через условие
3. В цикле проходимся по нашей строке
4. Получаем цифру через преобразование. Прием: `int(byte - '0')`. Таким образом получаем нужную цифру из байта
5. Проверяем не выходит ли наш результат за пределы диапазона 32-разрядного целого числа
6. Если встречается не число (0-9), то завершаем цикл и выводим результат со знаком