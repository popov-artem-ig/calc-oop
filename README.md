# calc-oop
Обратная польская запись реализована как тип ReversePolishNotation
- Calc() (float64, error)
- ToString() (string, error)
- IsEmpty() bool

Конвертация происходит сразу при создании
Вычисление при вызове метода Calc

Стек и Очередь так же реализованы отдельными типами

ReversePolishNotation, Stack, Queue автоматически реализуют интерфейс ICustomType, так как реализуют IsEmpty и ToString 
могут быть переданы параметрами в DebugLog(data ICustomType)
