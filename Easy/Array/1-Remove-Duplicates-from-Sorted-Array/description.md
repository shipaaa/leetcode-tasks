# Task Description

## RUS

При наличии целочисленного массива `nums`, отсортированного в порядке неубывания, удалите дубликаты на месте, чтобы каждый уникальный элемент отображался только один раз. Относительный порядок элементов должен оставаться неизменным. Затем верните количество уникальных элементов в `nums`.

Предположим, что количество уникальных элементов `nums` равно `k`, для того, чтобы их приняли, вам нужно выполнить следующие действия:
- Измените массив `nums` таким образом, чтобы первые `k` элементов `nums` содержали уникальные элементы в том порядке, в котором они изначально присутствовали в `nums`. Остальные элементы nums не важны, так же как и размер `nums`.
- Верните `k`.

## ENG

Given an integer array `nums` sorted in **non-decreasing order**, remove the duplicates [**in-place**](https://en.wikipedia.org/wiki/In-place_algorithm) such that each unique element appears only **once**. The relative order of the elements should be kept the **same**. Then return the number of unique elements in `nums`.

Consider the number of unique elements of `nums` to be `k`, to get accepted, you need to do the following things:
- Change the array `nums` such that the first `k` elements of `nums` contain the unique elements in the order they were present in `nums` initially. The remaining elements of `nums` are not important as well as the size of `nums`.
- Return `k`.

## Проверка/Testing

```go
int[] nums = [...]; // Input array
int[] expectedNums = [...]; // The expected answer with correct length

int k = removeDuplicates(nums); // Calls your implementation

assert k == expectedNums.length;
for (int i = 0; i < k; i++) {
    assert nums[i] == expectedNums[i];
}
```

## Объяснение задачи
1. Мы используем переменную `k`, чтобы отслеживать количество уникальных чисел. Изначально она равна 1, потому что первый элемент массива всегда уникален
2. Мы проходимся по массиву, начиная со второго элемента и сравниваем его с последним уникальным элементом, который мы уже нашли (это элемент на позиции `k-1`)
3. Если текущий элемент отличается от последнего уникального элемента, это новый уникальный элемент. Мы помещаем его в массив на позицию `k` и увеличиваем `k` на 1
4. После завершения обхода массива `k` будет содержать количество уникальных элементов. Массив будет изменён так, что первые `k` элементов будут уникальными