Задача 1 
В першій і другій задачі я використав цикл для того, щоб не писати код до кожного варіанту(тобто 8,16,32 біт і так далі), а зробив межі від 8 до 4096 з кроком i*2, Для великих чисел використовував тип big.Int, Створив змінну результату типу big.Int та використав в ній метод Exp який дає нам можливість піднести 2 до степеня i(до вибраної кількості бітів), не знайшов для себе інший спосіб так як потрібно було результат помістити в тип big.Int. І через пробіл виводиться результат(кількість можливих варіантів до 8,16 бітів і так далі) до всіх варіантів які були вказані в завданні.
Задача 2
З важливого в другій задачі для генерації рандомного числа я використовував пакет crypto/rand, усе що було потрібно це об'явити змінну maxNum яка зберігає в собі максимальне можливе значення під виділенні біти, тобто якщо для прикладу 8 біт то це рахується як 2^8 - 1, а пакет crypto/rand дає нам можливість використати функцію rand.Int в якій ми вказуємо максимальне можливе значення числа під відповідну кількість бітів, і воно буде вибирати рандомне число від 0 до макс можливого числа і що важливо значення буде мати тип big.Int, тобто ми зможемо запихнути в змінну з цим типом будь-яке число.
Задача 3
В цій задачі я створив функцію брутфорса,суть функції полягає в тому що вона приймає на вхід значення яке ми брутфорсимо, і в ній є бескінечний цикл який закінчиться коли ми знайдемо потрібне нам значення, в самому циклі  просто перевіряється через метод Cmp чи рівні два числа big.Int якщо ні то за допомогою методу Add додаю одиницю до числа n яке використовується для перебору . Також функція повертає значення яке ми шукали і через прінт виводить затрачений час, і в функції main в кінці є виклик цієї функції з оголошеною тестової змінною значення якої я подав на вхідні дані функції. Одразу хочу вибачитись якщо написав забагато або замало, просто не маю уявлення коли саме потрібно пояснювати свої функції і чи потрібно було це робити в мому випадку, але в будь-якому випадку лишнім не буде.
