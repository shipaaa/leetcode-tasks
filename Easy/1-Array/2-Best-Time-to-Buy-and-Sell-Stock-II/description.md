# Task Description. [Best Time to Buy and Sell Stock II](https://leetcode.com/explore/interview/card/top-interview-questions-easy/92/array/564/)

## RUS

Вам дан целочисленный массив `prices`, где `prices[i]` — это цена данной акции на i-ый день.

В любой день вы можете принять решение о покупке и/или продаже акции. Вы можете владеть не более чем одной акцией в любое время. Однако вы можете купить их, а затем сразу же продать в тот же день.

Найдите и верните максимальную прибыль, которую сможете получить.

## ENG

You are given an integer array `prices` where `prices[i]` is the price of a given stock on the `i-th` day.

On each day, you may decide to buy and/or sell the stock. You can only hold **at most one** share of the stock at any time. However, you can buy it then immediately sell it on the **same day**.

Find and return _the **maximum** profit you can achieve_.
## Объяснение
1. Создаем переменную `lastPrice` и присваиваем ей значение `цена акции за предыдущий день`.
2. Проходимся в цикле по ценам всех дней до тех пор, пока наш массив не закончится.
3. Проверяем условие если нынешняя цена больше цены за предыдущий день, то вычитаем нынешнюю цену акции и цены за предыдущий день и записываем в общий профит (переменная `profit`)
4. В конце цикла (конец каждого дня) задаем переменной `lastPrice` сегодняшнюю цену `price` и переходим к следующему дню.